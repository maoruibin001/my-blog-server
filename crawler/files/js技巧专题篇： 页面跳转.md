# js技巧专题篇： 页面跳转 

本篇主要介绍网页上常见的页面跳转技术。页面跳转有几种方式，比较常用的是window.location.href,window.location.replace,window.open，当然还有目前比较火的很多框架都采用的无刷新页面跳转技术window.history.pushState和window.history.replaceState。这些我都不讲^_^，我这里讲得是meta的一个相关配置。我相信，很多朋友看见实现的页面时会非常面熟，特别是男性！

以下是相关代码实现：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <meta http-equiv="refresh" content="5;url=https://www.baidu.com"/>
    <title></title>
    <style>
        span {
            color: red;
            padding: 5px 15px;
            background: #cccccc;
        }
        button {
            padding: 10px;
            display: inline-block;
            vertical-align: top;
            border-radius: 4px;
            outline: none;
        }
    </style>
</head>
<body>
<h1>对不起您浏览的页面已改变，<span>5</span> 秒后自动为您跳转... <button>手动跳转</button></h1>
<script>
    var span = document.querySelector('span'),
            btn = document.querySelector('button');
    var selfTimer = (function(){
        var i = 5;
        return function(){
            span.innerHTML = --i;
            if (i == 0) {
                clearInterval(timer);
            }
        }
    })()
   timer = setInterval(selfTimer, 1000);

    btn.onclick = function() {
        window.location.hash = 'https://www.baidu.com';
    }
    </script>
</body>
</html>```

哈哈哈，如果有人运行这段代码就懂了^_^，就差一个我们的新网址是1024，懂的自然懂。

