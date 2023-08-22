# weex踩坑笔记 

1、weex App版本不能直接使用vue组件的标签形式，例如：

```javascript<TabBar ></TabBar >

只能使用：<component :is="currentView2" :barItems="barItems">
这种方式。```

2、weex中屏幕高度不能直接使用weex.config.env.deviceHeight而要取：
750 / weex.config.env.deviceWidth * weex.config.env.deviceHeight
因为weex默认的屏幕宽度为750，并且所有情况都使用默认750.

3、weex在web平台上的this.$refs.xxx找到的是一个vue实例，如果要找元素需要this.$refs.xxx.$el才行。

4、weex input中清除value之后直接调用focus不行，必须在setTimeout中调用才能达到想要的效果。

5、weex中想要实现input宽度为容器的百分比，只能在input元素上加上width，如下：
input width="100%" ref="input"如果在style中设置width在app中无效。

6、weex中想要app和web页面都同时获取到value，可以在change或者input事件中使用v-model或者event.value。 inputText(event) {
this.showX = !!event.value;
}

7、bug：weex app中不能清除input内容。

8、weex app中切换清除borderWidth的时候不能设置为none，必须设置

9、weex浏览器触发touchend事件有异常，一般要两次才能成功触发一次，所以防重复限制锁最好在touchstart上加，不要加在touchend上。

10、weex的touch事件在不同平台下表现不一致：想要获取水平偏差，如下：

```javascriptgetXOffset(event) {
    return PLATFORM === 'Web' ? event.touches[0].clientX : event.changedTouches[0].screenX;
},```

