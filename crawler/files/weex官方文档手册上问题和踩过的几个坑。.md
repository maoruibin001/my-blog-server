# weex官方文档手册上问题和踩过的几个坑。 

1、list下拉刷新bug（web端的时候根本不会正常触发loadmore事件）

2、image注意写宽高，input注意字体大小（特别是input在web端表现和app端完全两样，有时候甚至没法正常输入）。

3、modal.toast  没有清空定时器的功能，注意做防重复处理，可以用setTimeout做一个同步更新状态。

4、web端使用picker需要做额外处理（需要在index.html中额外引入：<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="../node_modules/@weex-project/weex-picker/js/build/index.js"></script>  ）。

```javascript<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="../node_modules/@weex-project/weex-picker/js/build/index.js"></script> ```

5、stream.fetch注意post请求的请求头问题(fetch请求需要特别注意content-type的类型，特别是application/json和application/x-www-form-urlencoded的数据格式是不一样的，不要用混淆了。并且put、delete请求不能发送application/json的数据类型)。

6、input没有v-model。

7、animation的position定位问题，在手机端不支持left、top这样的位移方式，可以用transform来做兼容，但切记web端必须加上单位才能正常执行，app端不用加上单位也能正常执行。

8.modal 中alert之类的duration问题（alert、confirm、prompt加上duration完全是没用的，应该是手册上存在的笔误）。

9、animation中styles中transform存在的问题（显示为一个对象，但只能接受一个字符串，多个动画中间用空格隔开，如：transform: 'translate(500px, 500px) rotate(3600deg)'）。

当然手册上还有些问题，当时没记下来，后再再去找又一时找不到，如果大家还有发现其他问题，可以在下面留言。

