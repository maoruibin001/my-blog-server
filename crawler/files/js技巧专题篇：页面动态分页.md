# js技巧专题篇：页面动态分页 

本篇主要介绍网页上常见的页面分页技术，用前端实现的方式更加省时省力（以前做过一个购物后台，用php实现的动态分页），这里只写了一种简单的实现方式，只是提供一种思路。

以下是相关html代码实现：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
    <style>
        td {
            border: 1px solid #ccc;
            height: 100px;
            width: 200px;
            text-align: center;
            font-size: 10pt;
            padding: 5px;
        }
    </style>
</head>
<body>
<table id="tablePaging" border="1">
    <tr>
        <td>第一页内容</td>
        <td>第一页内容</td>
    </tr>
    <tr>
        <td>第一页内容</td>
        <td>第一页内容</td>
    </tr>
</table>
<p class="paging">
    总页数 <span id="allPage">1</span>
    当前页数 <span id="currentPage">1</span>
    <input type="button" value="上一页" id="prevPaging"/>
    <input type="button" value="下一页" id="nextPaging"/>
</p>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/changePage.js"></script>
<script>
    tablePaging({
        'tablePaging': document.getElementById('tablePaging'),
        'currentPage': document.getElementById('currentPage'),
        'allPage': document.getElementById('allPage'),
        'nextPaging': document.getElementById('nextPaging'),
        'prevPaging': document.getElementById('prevPaging')
    })
</script>
</body>
</html>```

相关js代码实现如下：

```javascript/**
 * Created by MAORUIBIN on 2016-03-30.
 */
;(function(window) {
    var win = window;
    var activeTr = function(table, num, tr) {
        if (!tr) {
            //删除操作
            var _num = table.rows[num];
            if (_num) {
                table.deleteRow(num);
                return true;
            }else {
                return false;
            }
        }else {
            //插入操作
            var row = table.insertRow(num),
                i = 0,
                len = tr.length;
            for (; i < len; ++i) {
                row.insertCell(i).innerHTML = tr[i];
            }
            return true;
        }
    };

    var currentPage = 1,
        table = null,
        currentPageUi = null,
        allPage = null;

    var updateUi = function() {
        tableUi();
        currentPageUi.innerHTML = currentPage;
            allPage.innerHTML = allPages;
        },
        getPageData = function() {
            return [
                ['第' + currentPage + '页内容','第' + currentPage + '页内容'],
                ['第' + currentPage + '页内容','第' + currentPage + '页内容']
            ]
        },
        allPages = 5,
        tablePaging = function(args) {
            table = args.tablePaging;
            currentPageUi = args.currentPage;
            allPage = args.allPage;

            nextPaging(args.nextPaging);
            prevPaging(args.prevPaging);

            updateUi();
        },
        nextPaging = function(ele) {
            ele.onclick = function() {
                currentPage ++;
                if (currentPage > allPages) {
                    currentPage = allPages;
                    return;
                }
                updateUi();
            }
        },
        prevPaging = function(ele) {
            ele.onclick = function() {
                currentPage --;
                if (currentPage < 1) {
                    currentPage = 1;
                    return
                }
                updateUi();
            }
        },
        tableUi = function() {
            var data = getPageData(),
                _datal = null,
                len = data.length,
                i = 0;
            for(; i < len; ++i) {
                activeTr(table, 0);
            }
            for(i = 0; i < len; ++i) {
                _datal = data[i];
                activeTr(table, 0, [
                    _datal[0],
                    _datal[1]
                ]);
            }
        };

    win.tablePaging = tablePaging;
})(window)```

