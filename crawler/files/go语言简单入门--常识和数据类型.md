# go语言简单入门--常识和数据类型 

go语言常识：

1、包导出机制：首字母大写，代表可以导出，否则不能导出，所以看到所有包的方法都是首字母大写的。

2、所有的导入都要放到文件最前面，这是解析机制导致的。

3、go关键字可以并发，但是执行依赖调度，如果只是在main顺序的新建goroutine没有触发调度，则goroutine里面的代码不会执行。



go语言数据类型

1、go语言数据类型

a) 基本类型、引用类型、聚合类型、接口类型。（其实可以更广泛的将这个类型分为基本类型和引用类型，

b) 如何区分数据类型：基本类型空值是当前类型，引用类型的空值是nil。一般情况下，基本类型可以比较（==）， 引用类型不可比较。）

2、基本类型（18个）

bool

byte

rune（存储unicode）

int/uint

int8/uint8

int16/uint16

int32/uint32

int64/uint64

float32

float64

complex64

complex128

string

a) Int: int8、 uint8、int16、uint16...，int需要注意的点是：

i. 基础的int类型在不同的操作系统里是不一样的，64位操作系统中是int64，32位操作系统中是int32.

ii. Byte是int8类型的同义词，rune是int32的同义词，前者是一个数据的原始类型，后者多用于描述unicode码点。

b) float32、float64

c) complex64、complex128

d) string

e) struct

f) bool

g) array [v]type{} 这里也可以用特殊表示法如[...]int{99: 1}

 

3、常见的引用类型

a) Slice：slice看上去和array类似，但是是array的引用即每一个slice都有一个底层数组。[]type{value1}

b) Interface

c) Map: map[keyType]valueType

d) channel

 

- 这里有几点小技巧：

这里有几点小技巧：

可以比较的数据类型才能作为一个map的键，所以map的键都是基本类型。

值类型存储在栈中，引用类型存储在堆或栈中。

无符号整数的应用场景大多是些二进制解码，加密或hash等。

Const类型的变量都是基本类型。

 

- String类型常用的包有下面四个：

String类型常用的包有下面四个：

strings：用于对string进行一些操作，例如count、prefex、contains等等

strconv: 用于string与其他类型的装换或者解析，如strconv.Itoa、strconv.ParseBool()等。

bytes:和strings类似，不过大多处理的是[]byte{}这种类型的字符串。

unicode: 处理unicode相关的东西。

 

需要记住：

通过fmt.Print之类的函数打印一个变量时，其实是调用当前变量的String方法，因 此只要这个变量包含String方法，就会重写原有的String方法，调用当前变量String 方法。 





