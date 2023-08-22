# http你不得不知道的那些事（七）--跨域请求 

之前有篇文章详细的讲述了浏览器的同源策略，并且附带的提了下跨域。本篇就详细的分享下跨域请求方案。

1、jsonp

JSONP(JSON with Padding)是JSON的一种“使用模式”，可用于解决主流浏览器的跨域数据访问的问题。由于同源策略，一般来说位于 server1.example.com 的网页无法与不是 server1.example.com的服务器沟通，而 HTML 的<script> 元素是一个例外。利用 <script> 元素的这个开放策略，网页可以得到从其他来源动态产生的 JSON 资料，而这种使用模式就是所谓的 JSONP。用 JSONP 抓到的资料并不是 JSON，而是任意的JavaScript，用 JavaScript 直译器执行而不是用 JSON 解析器解析（来自百度百科）。其实简单点讲就是利用script标签的src属性不会被同源策略限制，并且请求下来的文件直接执行的策略。具体实现为：

我们可以先在当前页类设置一个回调函数，然后创建一个script标签，在src中传入我们准备好的回调函数。服务器接收到请求，提取出回调函数，然后把想要传给客户端的内容塞入回到函数的参数中，返回客户端。客户端会直接执行代码，调用回调函数，回调函数参数就是从服务器返回的内容。

jsonp是最常用，也是比较方便的跨域解决方案。如果要完整的实现一个jsonp方案，必须服务器端和客户端同时配合，少了一方，请求就会没有意义。

2、带src属性的标签跨域

基本所有的带有src属性的标签都能跨域请求资源，但是由于请求资源不会向script标签一样直接执行，因此就需要和一些其他方案组合起来才能实现完整的跨域通信。常见的如跨域iframe的postMessage，url截取等。特别举个实例，在当前页内（第一个页面）跨域嵌入一个评论页面（第二个页面），需要嵌入的页面高度随评论高度增加而增加。这里以上两种方案都可以实现，postMessage很好理解。采用url截取的话，必须引入第三个页面，第三个页面必须和我们当前页面同源且必须嵌套在待嵌入的评论页面中。因为第三个页面和第一个页面同源，这样就可以通过parent.parent.xxx访问第一个页面的文档，是不是很巧妙~

3、websocket跨域

websocket和http不同，浏览器没有对它进行同源策略的限制。是不是感觉不可思议，不能理解，这就要同父不同命~。这些东西反正都是大牛捣鼓出来的额，他们想咋弄就咋弄，我们是菜鸟也就只能学习了。具体实现很简单，只需要建一个websocket服务器，然后大家都和websocket服务器交流就可以了。我这里就提一下如何发起websocket请求，此处又有一个奇葩的地方，websocket的握手采用的是http协议！！请求头信息中加入这两条就行了，Connection:Upgrade
Upgrade:websocket
其他的就是正常http请求，如下：
GET / HTTP/1.1
Connection:Upgrade
Host:127.0.0.1:8088
Origin:null
Sec-WebSocket-Extensions:x-webkit-deflate-frame
Sec-WebSocket-Key:puVOuWb7rel6z2AVZBKnfw==
Sec-WebSocket-Version:13
Upgrade:websocket
不想多说，虽然很奇怪，但是我是码农，我不会思考，我只会抄~~

4、document.domin跨域

这个跨域方案只适合用在父域和子域之间跨域，具体是将子域domin.name设置成父域，然后浏览器就会假装他们两个是同域了~~

5、CORS跨域

CORS是一个W3C标准，全称是"跨域资源共享"（Cross-origin resource sharing）。
它允许浏览器向跨源服务器，发出XMLHttpRequest请求，从而克服了AJAX只能同源使用的限制。以express为例：

var express = require('express');
var app = express();
//设置跨域访问
app.all('', function(req, res, next) {
res.header("Access-Control-Allow-Origin", "");
res.header("Access-Control-Allow-Headers", "X-Requested-With");
res.header("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS");
res.header("X-Powered-By",' 3.2.1')
res.header("Content-Type", "application/json;charset=utf-8");
next();
});  

app.get('/auth/:id/:password', function(req, res) {
res.send({id:req.params.id, name: req.params.password});
});  

app.listen(3000);
最关键的设置响应头"Access-Control-Allow-Origin", "*"。

6、window.name跨域
window.name跨域的核心思想是同一个tab页即使页面跳转window.name也不会改变，例如你在a.html中设置了window.name='maoruibin'，并且<a href='b.hmtl'>b.html</a> 点击调转后在b.html中查看window.name也是显示也是“maoruibin”。因此window.name结合iframe跨域通信也很好理解了，如下：
a.html:

```javascript<!doctype html>
<html>
<head>
    <meta charset="UTF-8">
    <title>127.0.0.1:8088</title>
</head>
<body>
<h1>127.0.0.1:8088</h1>
<script>
  function test(){
    var obj = document.getElementById("iframe");
    obj.onload = function(){
      var message = obj.contentWindow.name;
      console.log(message);
    }
    obj.src = "about:blank";
  }
</script>
<iframe style="display:none;" id="iframe" class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="http://localhost:8089" onload="test()"></iframe>
<iframe  class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="http://127.0.0.1:8089/b.html" ></iframe>
</body>
</html>```

b.html:

```javascript<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>window.name</title>
</head>
<body>
<h1>localhost:8089</h1>
<script>
  //todo
  window.name = "This is message!";
</script>
</body>
</html>```

此时我们打开a.html既能在控制台看到b.html跨域传来的信息，也能看到b.html的内容，如下：

这里有个注意点：

about:blank，javascript: 和 data: 中的内容，继承了载入他们的页面的源。
好了，以上就是所有跨域的内容。window.name跨域可以自己建一个小服务器，自己动手试一下，这样才能记得住。thx~

后语：http相关的内容，比较多，我分享的也都是九牛一毛，不过能力有限，这篇之后http相关的东西，暂时就不会更新了。以后大概会写点前端性能相关的东西，也可能做点websocket的内容分享，到时候再看吧，拜拜~

