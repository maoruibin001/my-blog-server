# js技巧专题篇：动态增删一行表格 

本篇介绍的是在表格中动态添加删除一行，这在购物车类的实现中非常常见。我采用的是表格的形式，因为大都在购物类呈列商品时都采用的表格（我想这个表格在前端最有用最好用的地方了）。并且由于表格布局曾经辉煌过，所以它有很多API接口，操作起来很方便。

以下是相关html代码实现：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<table>
    <tr>
        <td>1</td>
        <td>2</td>
        <td>3</td>
    </tr>
    <tr>
        <td>4</td>
        <td>5</td>
        <td>6</td>
    </tr>
    <tr>
        <td>7</td>
        <td>8</td>
        <td>9</td>
        <td>10</td>
    </tr>
</table>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/activeTr.js"></script>
<script>
    var table = document.querySelector('table');
    activeTr(table ,2);
    activeTr(table, 2, ['新增单元格1', '新增单元格2']);
</script>
</body>
</html>```

下面是js代码实现：

```javascript/**
 * Created by MAORUIBIN on 2016-03-30.
 */
(function(window) {
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

    win.activeTr = activeTr;
})(window)```

