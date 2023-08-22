# go语言简单入门--包、测试和低级编程 

go语言常用的包：

strings、strconv、bytes、unicode

fmt、log、errors

math

net、http

json、html

io、bufio

time、flag、sort

sync

注意点：

new(Type) *Type 。new一个类型返回的是该类型的指针。

json.Marshal能序列化的字段必须首字母大写（代表能导出），想要输出为小写则必须在后面打上json标签如：Name string `json: "name"`

json.Unmarshal()第一个接受[]byte{}类型变量，第二个参数必须传入指针（如果不是指针，则解码之后的值不会更新,且会返回错误）。该解码返回值为error类型，表示是否成功。



go测试：

前提：

1、在一个文件夹下必须有一个[name]_test的文件如getName_test

2、测试函数必须是Test开头如TestGetName



go test运行的一些参数

go test -v输出测试用例名称和运行时间

go test -run支持正则，可以筛选运行测试的文件 。如go test -run="Hello | world"



查看一个包中的build文件和测试文件指令

go list -f={{.GoTestFiles}} net

go list -f={{.GoFiles}} fmt

go list -f={{.XGoTestFiles}} os



低级编程



go语言中提供了对低级编程的支--unsafe包，大多数时候我们不使用这个包。go中内置模块如os、system、io、net等大量使用这个包。

偶尔会用到的函数如unsafe.SizeOf()返回变量在内存中占用的字节数。unsafe.OffsetOf()获取成员变量相对于结构体起始位置的偏移量，在某些编解码时可能会用到。





