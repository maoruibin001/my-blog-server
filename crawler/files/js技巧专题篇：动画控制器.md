# js技巧专题篇：动画控制器 

这篇文章是整个技巧篇第一阶段的终结，也是大家比较关心的特效动画实现一个控制器。相对来说这篇代码会比较多一些，涉及的知识点相对来说更全面，懂一些设计模式的朋友可能更容易理解看似复杂实现背后的思路。当然，这篇文章只是提供一个想法，一个思考问题的方式，而不是为了简单方便。这个控制器在实现少量动画的时候其实就是大材小用，得不偿失，这也是很多设计模式必然会面临的问题，希望能对需要的朋友一点借鉴。

相关html代码如下：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
    <style>
        img {
            position: absolute;
        }
    </style>
</head>
<body>
<img class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="img/yellowStar.png" alt="" id="myTest" style="left: 0;top: 0;"/>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/animationController.js"></script>
<script>
    var myTest = document.getElementById('myTest');
    myTest.onclick = function() {
        new animationController({
            'context': this,
            'effect': 'linear',
            'time': 500,
            'startCss': {
                'left': this.style.left,
                'top': this.style.top
            },
            'endCss': {
                'left': 200,
                'top': 200
            },
            'callback': function() {
                alert('动画结束');
            }
        }).init();
    }
</script>
</body>
</html>```

相关js代码：

```javascript/**
 * Created by MAORUIBIN on 2016-04-06.
 */
;(function(window) {
    var win = window,
        _queue = [],
        _baseUID = 0,
        _updateTime = 1000,
        _ID = -1,
        isTicking = false;
    /*options参数
    * context--被操作元素的上下文
    * effect--动画效果算法
    * time--效果的持续事件
    * startCss--元素的起始偏移量
    * endCss--元素的结束偏移量
    * */
    var animationController = function(options) {
        this.context = options;
    }
    animationController.prototype = {
        init: function() {
            this.start(this.context);
        },
        start: function(options) {
            options && this.pushQueue(options);
            if (isTicking || _queue.length === 0) return false;
            this.tick();
            return true;
        },
        stop: function() {
            clearInterval(_ID);
            isTicking = false;
        },
        tick: function() {//动画检测
            var self = this;
            isTicking = true;
            _ID = setInterval(function() {
                if (_queue.length === 0) {
                    self.stop()
                } else {
                    for (var i = 0, len = _queue.length; i < len; ++i) {
                        _queue[i] && self.run(_queue[i], i);
                    }
                }
            }, _updateTime);
        },
        run: function(_options, i) {
            var now = this.now(),
                st = _options.startTime,
                timing = _options.time,
                e = _options.context,
                t = st + timing,
                name = _options.name,
                tPos = _options.value,
                sPos = _options.startValue,
                effect = _options.effect,
                scale = 1;

            if (now >= t) {
                _queue[i] = null;
                this.delQueue();
            } else {
                _tPos = this.effect({
                    e: e,
                    timing: timing,
                    now: now,
                    st: st,
                    sPos: sPos,
                    tPos: tPos
                }, effect);
            }
            e.style[name] = name == 'zIndex' ? tPos : tPos + 'px';
            this.callback(_options.callback, _options.uid);
        },
        effect: function(_options, effect) {
            effect = effect || 'liear';
            var _effect = {
                'linear': function(_options) {
                    var scale = (_options.now - _options.st) / _options.timing,
                        tPos = _options.sPos + (_options.tPos - _options.sPos) * scale;
                    return tPos;
                }
            }
            return _effect[effect](_options);
        },
        callback: function(callback, u) {
            isCallback = true;
            for (var i = 0, len =  _queue.length; i < len; ++i) {
                if (_queue[i].uid == u) isCallback = false;
            }
            isCallback && callback && callback();
        },
        pushQueue: function(options) {
            var ctx = options.context,
                t = options.time || 1000,
                callback = options.callback,
                effect = options.effect,
                startCss = options.startCss,
                endCss = options.endCss,
                u = this.setUID(ctx);
            for (var name in endCss) {
                _queue.push({
                    'context': ctx,
                    'time': t,
                    'name': name,
                    'value': parseInt(endCss[name], 10),
                    'startValue': parseInt((startCss[name]) || 0),
                    'effect': effect,
                    'uid': u,
                    'callback': callback,
                    'startTime': this.now()
                })
            }
        },
        delQueue: function() {
            for (var i = 0, len = _queue.length; i < len; ++i) {
                if (_queue[i] === null) _queue.splice(i, 1);
            }
        },
        now: function() {
            return new Date().getTime();
        },
        getUID: function(_e) {
            return _e.getAttribute('UID');
        },
        setUID: function(_e, _v) {
            var u = this.getUID(_e);
            if (u) return u;
            u = _v || _baseUID ++;
            _e.setAttribute('UID', u);
            return u;
        }
    }

    win.animationController = animationController;
 })(window)```

