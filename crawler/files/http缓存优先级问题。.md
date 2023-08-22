# http缓存优先级问题。 

昨天公司运维来找测试，希望他们注意响应中的缓存字段。我在旁边听着，觉得他们讲得很有问题，就插了两句，和他们说了下公司目前官网图片缓存策略的问题。然后就让我写个小文档发出来，好吧，谁叫我多嘴。这篇文章就是自己整理的一个小文档，详细的描述了http缓存及其优先级。

HTTP缓存

主要有两种缓存：强缓存和对比缓存（也叫协商缓存）。

强缓存：只要请求了一次，在有效时间内，不会再请求服务器（请求都不会发起），直接从浏览器本地缓存中获取资源。主要字段有（expires：date（过期日期）、cache-control： max-age=time（毫秒数，多久之后过期）                       |no-cache|no-store）。如果expires和cache-control同时存在，cache-control会覆盖expires。建议两个都写，cache-control是http1.1的头字段，expires是http1.0的头字段，都写兼容会好点。

对比缓存：无论是否变化，是否过期都会发起请求，如果内容没过期，直接返回304，从浏览器缓存中拉取文件，否则直接返回更新后的内容。主要字             段有：

1、服务器端返回字段 Etag： xxxx (一般为md5值) 对应客户端匹              配字段为， If-None-Match: w/xxx(xxx的值和上面的etag的xxx一样则返                回304，否则返回修改后的资源)。

2、服务器端返回字段：Last-Modifieddate(日期),对应客户端匹配字段If-Modified-Since:date（如果服务器date小于等于客户端请求date则返回304，否则返回修改后的资源））。

同时存在各种缓存头时，各缓存头优先级及生效情况:

1、强缓存和对比缓存同时存在，如果强缓存还在生效期则强制缓存覆盖对比缓存，对比缓存不生效；如果强缓存不在有效期，对比缓存生效。即：强缓存优先级 > 对比缓存优先级

2、强缓存expires和cache-control同时存在时，则cache-control会覆盖expires，expires无论有没有过期，都无效。 即：cache-control优先级 > expires优先级。

3、对比缓存Etag和Last-Modified同时存在时，则Etag会覆盖Last-Modified，Last-Modified不会生效。即：ETag优先级 > Last-Modified优先级。

当然还有一种缓存pragma，和cache-control类似，前者是http1.0内容后者是http1.1内容，并且pragma优先级 > cache-control优先级，不过前者目前基本不使用。

针对我们当前的项目，由于css和js在打包时加了md5值，建议直接使用强缓存，并且expires和cache-control同时使用，建议设置时长为7天较为妥当。图片文件由于没有加md5值，建议采用对比缓存，html文件也建议采用对比缓存。

ps: 当我们不设置cache-control，只设置对比缓存，在不同浏览器下会有不同的表现。chrome会直接从本地缓存获取，其他会请求服务器返回304.这时候有两种方式让他们的响应一致。1、设置cache-control: public, max-age=0;记住，这里的public是关键。因为默认值是private，表示其他代理都不要缓存，只有服务器缓存，而max-age又为0，所以每次都会发起200的请求。设置public的意思就是允许其他各级代理缓存资源，因此如果资源没改变会返回304。 2、直接设置max-age=1000。即是一秒之后内容过期，目的是触发浏览器缓存。也能达到想要304的效果。

