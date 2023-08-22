# indexedDb使用及遇到的问题 

indexedDb作为前端存储大数据首选方案，现在已经越来越流行。本文主要简单介绍indexeddb、以及遇到的坑和解决方案。

### 简介

indexedDb（以下简称IDB），操作还是比较复杂的，主要涉及到数据库、对象仓库、索引、事物、指针、主键集合、操作请求。

IDB存储空间的大小，主要由IDB存储位置的磁盘大小确定，官网上说最多是剩余空间50%，一般是符合2^n定律。chrome存储文件位置为：C:\Users[用户名]\AppData\Local\Google\Chrome\User Data\Default\IndexedDB[http域名]，之前官网看到和域名相关（组），经验证并没有。

IDB主要操作都为异步操作，都涉及到成功、中断、失败回调。最佳的处理方式是将操作封装成promise，然后用async、await方式调用。

### 坑及解决方案

1、IDB open时，如果数据库不存在，则新建，如果存在则打开。新建或者升级版本时，会触发onupgradeneeded事件，之后会触发onsuccess事件。正常打开时，只会触发onsuccess事件。

2、IDB所有的操作都需要通过事物进行，每个事物只能进行一次操作。事物存在几种状态，激活、过期、失败、成功、中断、完成。这里有个比较坑的地方是，首次创建数据库时，会触发onupgradeneeded事件，此时事物只能使用event.target.transaction，而不能通过db去获取。
事物操作数据有两种方式

3、IDB的delete操作会删掉记录，该记录占用的存储空间可能不会马上释放，需要等待一定的时机才会删除，按照学术点的话就是先标记删除，到特定时机再删除。在stackoverflow上看到有人说是需要再增加4M才会触发删除，经验证，并不是。回收也会出现不彻底的情况，可能会导致最终数据不多，单占用空间持续增大的情况。删除时最好不要用指针遍历一个一个删，最好是用主键集合的方式删除，这里就涉及到IDBKeyRange对象。

4、IDB数据库存满之后会报QuotaExceededError错误，此时删内容之后由于回收机制的限制，会出现删除内容，但是空间不释放的bug。此处大多处理方式是删库，删库也不能直接删，需要先close才能删除，否则会报错。应对这种情况较好的处理方式是：1、尽量不让库满，采用的方式大多有两种，限制条数和限制存储时间。2、库满之后取出最近的部分条目，删除，然后再将内容写入新库。这就注定了IDB不适合存储需要完整内容的条目。

5、webkit内核可以通过navigator.webkitTemporaryStorage.queryUsageAndQuota查看总存储和使用的存储。具体方法为：
navigator.webkitTemporaryStorage.queryUsageAndQuota(
(usedBytes, grantedBytes) => {
// do something
},
e => {
reject(e)
}
)
如果在chrome浏览器中可以通过控制台=》application=>clear storage查看。

6、IDB 在有些非正常情况下会报QuotaExceededError错误，这种时候在某些特定情况下会出现db获取不到的情况（而通过db获取transaction却能获取到），此时直接关闭db会出错。

### 参考

IDB基本操作的网站建议参考：http://www.ruanyifeng.com/blog/2018/07/indexeddb.html

IDB官方文档建议参考：https://developer.mozilla.org/zh-CN/docs/Web/API/IndexedDB_API

后续可能会有更多的分享，如果你有问题也可以在下面留言。

