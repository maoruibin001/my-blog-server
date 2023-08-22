# js技巧专题篇：表单联动 

这篇文章要分享的一直技巧是表单联动，我们采用普适的方法写一个表单联动的函数库，当然和以前一样，我只是完成了其中的一部分，还有很大的扩展空间。有需要的朋友可以自行扩展，主要看气质！

相关html代码如下：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>表单联动</title>
</head>
<body>
<h1>表单联动</h1>
<form action="">
    请选择城市：
    <select name="parent" id="parent">
    </select>
    <br/>
    请选择区县：
    <select name="child" id="child">
        <option value="">请选择</option>
    </select>
</form>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/linkage.js"></script>
<script>
    var parent = document.getElementById('parent'),
            parentData = [
                {code: 0, value: '请选择'},
                {code: 1, value: '北京'},
                {code: 2, value: '上海'},
                {code: 3, value: '深圳'}
            ],
        child = document.getElementById('child'),
            childData = [
                ['请选择'],
                ['海淀区', '朝阳区', '丰台区', '石景山区', '通州区', '顺义区'],
                ['黄埔区', '浦东区 ', '长宁区', '宝山区 ', '杨浦区', '嘉定区','虹口区'],
                ['罗湖区', '福田区 ', '南山区', '保安新区 ', '光明新区']
            ];
    linkage(parent, parentData, child, childData);
</script>
</body>
</html>```

相关js代码如下：

```javascript/**
 * Created by MAORUIBIN on 2016-04-01.
 */
(function(window) {
    var win = window;
    var linkage = function(parent, parentData, child, childData) {
        _render(parent,parentData);
        parent.onchange = function() {
            var _value = this.value;
            _render(child, childData[_value])
        }
    }

    var _render = function(ele, data) {
        var opts = ele.querySelectorAll('option');
        for(var i = 0, len = opts.length; i < len; ++i) {
            ele.removeChild(opts[i]);
        }
        var frag = document.createDocumentFragment();
        for (var i = 0, len = data.length; i < len; ++i) {
            if (typeof data[0] === 'object') {
                var opt  = new Option(data[i].value, data[i].code);
            }else {
                var opt = new Option(data[i], i);
            }
            frag.appendChild(opt);
        }
        ele.appendChild(frag);
    }

    win.linkage = linkage;
})(window)```

