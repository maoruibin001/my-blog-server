const axios = require('axios');
const cheerio = require('cheerio');
const fs = require('fs')
const path = require('path')

function convertToMarkdown($, element) {
    let markdown = '';
  
    $(element).children().each((index, element) => {
      const tagName = $(element).get(0).tagName.toLowerCase();
  
      switch (tagName) {
        case 'h1':
          markdown += `# ${$(element).text()}\n\n`;
          break;
        case 'h2':
          markdown += `## ${$(element).text()}\n\n`;
          break;
        case 'h3':
          markdown += `### ${$(element).text()}\n\n`;
          break;
        case 'ul':
          $(element).find('li').each((index, liElement) => {
            markdown += `- ${$(liElement).text()}\n`;
          });
          markdown += '\n';
          break;
        case 'p':
          markdown += `${$(element).text()}\n\n`;
          break;
        default:
          console.log(`Unsupported element: ${tagName}`);
          break;
      }
  
      if ($(element).children().length > 0) {
        markdown += convertToMarkdown($, element);
      }
    });
  
    return markdown;
  }

async function fetchRuanYifengBlog() {
  try {
    const response = await axios.get('http://www.ruanyifeng.com/blog');
    const html = response.data;

    const $ = cheerio.load(html);

    // 获取所有 '继续阅读全文' 链接
    const readMoreLinks = $('.asset-more-link a');

    // console.log('readMoreLinks: ', readMoreLinks)

    // 遍历每个 '继续阅读全文' 链接
    for (let i = 0; i < readMoreLinks.length; i++) {
      const readMoreLink = readMoreLinks[i];
      const pageUrl = $(readMoreLink).attr('href');
      // 发送请求获取继续阅读的内容
      const pageResponse = await axios.get(pageUrl);
      const pageHtml = pageResponse.data;

      const $page = cheerio.load(pageHtml);

      const markdown = convertToMarkdown($page, $page.root());
      console.log(markdown)

    //   // 提取内容
    //   const content = $page('#main-content').text().trim();
    //   console.log(content);

      const title = $page('#page-title').text().trim();
      const fileName = `../files/${title}.md`;
      const filePath = path.join(__dirname, fileName);
      fs.writeFileSync(filePath, markdown);

      
      // 可以在这里对内容进行处理，例如保存到文件或进行其他数据提取操作
    }

  } catch (error) {
    console.error(error);
  }
}

fetchRuanYifengBlog();