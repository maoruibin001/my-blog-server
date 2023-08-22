# js技巧专题篇：表单常用的密码强度检验 

今天忽然想做一个技巧分享专题，分享一些网页中常见的一些元素的实现。此处采用纯手写的方式，将不借助于任何插件或者函数库。当然很多地方的实现并不完善，仅能提供一个我想到的思路，希望能对大家有所帮助，这是我写作这些文章的初衷。

```javascript/**
 * Created MAORUIBIN on 2016-03-29.
 */
(function(win){
    var showStrength = function(_this, showWrap) {
        showWrap.style.fontFamily = 'Microsoft Yahei';
        var oValue = _this.value,
            len = oValue.length,
            strengthAll = 0;
        var color = ['red', 'orange', 'green'],
            strStrength = ['密码长度不得小于6', '密码强度为：初级','密码强度为：中级','密码强度为：高级'];
        var strength = function(str) {
            var code = str.charCodeAt(0);
            if (code >= 48 && code <= 57) {
                return 1;
            }else if (code >= 97 && code <= 122) {
                return 2;
            }else {
                return 3;
            }
        }
        if (len < 6) {
            showWrap.innerHTML = strStrength[0];
            showWrap.style.color = color[0];
        }else {
            for (var i = 0; i < len; ++i) {
                strengthAll += strength(oValue[i]);
            }
            if (strengthAll < 10) {
                showWrap.innerHTML = strStrength[1];
                showWrap.style.color = color[0];
            }else if (strengthAll >= 10 && strengthAll < 16) {
                showWrap.innerHTML = strStrength[2];
                showWrap.style.color = color[1];
            }else {
                showWrap.innerHTML = strStrength[3];
                showWrap.style.color = color[2];
            }
        }
    }
    win.showStrength = showStrength;
})(window)
```

这是js的实现。在html中我们需要这样引用：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>表单常用的密码强度检验</title>

</head>
<body>
<input type="text" id = 'oInput'/><span id = 'oSpan'></span>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/strength.js"></script>
<script>
    var input = document.getElementById('oInput');
    var span = document.getElementById('oSpan');
    input.onkeyup = function() {
        var self = this;
        showStrength(self, span);
    }
</script>
</body>
</html>```

