# js技巧专题篇：五星好评。 

这篇文章主要介绍的是五星好评的实现。星评在各大购物网站和服务型网站普遍存在，一个好的星评实现可以让用户体验更好。这篇文章采用最普遍的实现方法，没有运用高端技巧，只是使用了一些比较巧妙的方法，希望能对喜欢的朋友有所帮助。

相关html代码如下：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>星级评价</title>
    <style>
        #votes div {
            display: inline-block;
            padding: 10px;
            background-image: url(img/emptyStar.png);
        }
    </style>
</head>
<body>
<div id="votes">
    <div></div>
    <div></div>
    <div></div>
    <div></div>
    <div></div>
</div>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/starVotes.js"></script>
<script>
    var voteBox = document.getElementById('votes');
    starVotes(voteBox, 'div');
</script>
</body>
</html>```

对应的js代码如下：

```javascript/**
 * Created by MAORUIBIN on 2016-04-05.
 */
(function(window) {
    var win = window;
    var getEleByType = function(eles, type) {
        var eleArr = [];
        for (var i = 0, len = eles.length; i < len; ++i) {
            if (eles[i].nodeName.replace('#').toLocaleLowerCase() === type) {
                eleArr.push(eles[i]);
            }
        }
        return eleArr;
    }
    var starVotes = function(box, starStr) {
        var starNum = 0,
            varmark = true,
            clicked = false;
        var _starArr = getEleByType(box.childNodes, starStr);
        for (var i = 0, len = _starArr.length; i < len; ++i) {
            _starArr[i].setAttribute('data-num', i);
            _starArr[i].onmouseover = function() {
                var _num = this.getAttribute('data-num');
                mark = true;
                _clearStar(_starArr);
                for (var j = 0; j <= _num; ++j) {
                    _starArr[j].style.backgroundImage = 'url(img/yellowStar.png)';
                }
            }
            _starArr[i].onmouseout = function() {
                var _num = this.getAttribute('data-num');
                if (!clicked) {
                    if (mark) {
                        for (var m = _num; m >= 0; --m) {
                            _starArr[m].style.backgroundImage= 'url(img/emptyStar.png)';
                        }
                    }
                }else {
                    _clearStar(_starArr);
                    for (var j = 0; j <= starNum; ++j) {
                        _starArr[j].style.backgroundImage = 'url(img/yellowStar.png)';
                    }
                }
            }
            _starArr[i].onclick = function() {
                var _num = this.getAttribute('data-num');
                mark = false;
                clicked = true;
                for (var j = 0; j <= _num; ++j) {
                    _starArr[j].style.backgroundImage = 'url(img/yellowStar.png)';
                }
                starNum = _num;
                console.log(starNum)
            };

        }

        var _clearStar = function(_starArr) {
            var len = _starArr.length;
            for (var l = len - 1; l >= 0; --l) {
                _starArr[l].style.backgroundImage= 'url(img/emptyStar.png)';
            }
        }
    };

    win.getEleByType = getEleByType;
    win.starVotes = starVotes;
})(window);```

