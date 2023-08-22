# js技巧专题篇：关键词过滤和多选限制 

这是上一篇密码强度检验的续集，关键词过滤涉及到关键词过滤。虽然关键词过滤大多是由后台来处理，但是前端如果直接处理掉，就会减轻后台的任务，从而降低后台压力。多选限制主要是在分类等可以多选单限制选择个数的情况，例如我们在慕课网上发表文章时，里面的标签可以最多选择3个这种情况。

相关html代码如下：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>js技巧专题篇：关键词过滤和多选限制</title>
</head>
<body>
<textarea name="key1" id="key1" cols="30" rows="10">
    我喜欢做一些有意义的事情，例如分享文章，帮助需要帮助的人，就像大家对性感美女的喜  欢一样。
</textarea>
<button>过滤</button>
<br/>
<br/>
<select name="sel" id="sel" multiple>
    <option value="JavaScript">JavaScript</option>
    <option value="Html/Css">html/css</option>
    <option value="Html5">Html5</option>
    <option value="C">C</option>
    <option value="C++">C++</option>
    <option value="Java">Java</option>
</select>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/keywordfilter.js"></script>
<script>
    var txt = document.getElementById('key1'),
            btn = document.querySelector('button'),
            opt = document.getElementsByTagName('option');
    btn.addEventListener('click', function() {
        keywordfilter(txt);
    })
    forbidcheck(opt);
</script>
</body>
</html>```

相关js代码如下：

```javascript/**
 * Created by MAORUIBINon 2016-03-29.
 */
;(function(win) {
    var keywordfilter = function(txt) {
        var keyword = ['性', /['喜']{1}.{0,5}[欢]{1}/g];
        for (var i = 0; i < keyword.length; ++i) {
            txt.value = txt.value.replace(keyword[i], '***')
        }
    };
    var forbidcheck = function(sel) {
        var selectNum = 0,
            limit = 3;
        if (!Array.isArray(sel)) {
            sel = [].slice.call(sel);
        }
        for (var i = 0, len = sel.length; i < len; ++i) {
            sel[i].onclick = function() {
                if (selectNum < limit) {
                    this.selected = true;
                    selectNum ++;
                }else {
                    this.selected = false;
                }
            }
        }
    }

    win.keywordfilter = keywordfilter;
    win.forbidcheck = forbidcheck;
})(window)```

当然，这个精选做得实在是太简陋，大家不必模仿，领会精神就行了^_^。

