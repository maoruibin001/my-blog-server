# go语言简单入门--函数和方法 

Len和cap方法的区别：

果新长度小于容量，将不会更换底层数组，否则更换。容量的用途是在数据拷贝和 内存申请的消耗与内存占用之前的权衡。



- panic函数可以引发运行时恐慌
- recover函数可以捕获运行时恐慌

panic函数可以引发运行时恐慌

recover函数可以捕获运行时恐慌

 

注意：偶尔会看到函数没有函数体，例如append， 是因为使用了其他语言实现 

Go闭包表示方法： func closer() func() return type {} 后面那个func()return type是返回值的类型。

 

Go语言是有块级作用域的，所以很多时候go可以在块级作用域中声明变量代替闭包的效果。特别注意这里的块级作用域值的是包含在{}中的代码，因此这里很多时候需要在代码块中再单独声明一个变量来保存外面语句的变量。（和js的let制造的块级作用域不同）



函数和方法表示并不一致。方法的表示如下：

func (v typeV) methodName(...args) return type {

//方法体

}

和函数表示有个明显的区别是多了个(v typeV)，这个的用处就是说明这是哪个类型的方法。

 

Go语言的实现基本也是基于方法来的，当我们说当前变量实现了某个类型，主要就是说当前变量有某个类型的方法，并没有严格的实现和继承之类的。



go语言类型判断可以使用的方式：

一、v.(int)

二、switch 语句中使用v.(type)

三、reflect.TypeOf()







reflect常用地方：

判断类型可以使用reflect.TypeOf()

遍历结构体

判断引用类型是否相等reflect.DeepEqual





