# go语言简单入门--关键字与defer 

关键字

- 程序申明： import package
- 程序实体声明和定义：chan、const、func、interface、map、struct、type和var
- 程序流程控制：go、select、break、case、continue、default、的份儿、else、fallthrough、for、goto、if、range、return和switch
- 并发关键字：go、chan和select

程序申明： import package

程序实体声明和定义：chan、const、func、interface、map、struct、type和var

程序流程控制：go、select、break、case、continue、default、的份儿、else、fallthrough、for、goto、if、range、return和switch

并发关键字：go、chan和select

defer函数执行规则：

1、当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会正在结束执行。

2、当执行外围函数中的return语句是，只有其中所有的延迟函数执行完毕后，外围函数才会真正的返回。

3、当外围函数中的代码引发运行时恐慌时，只有所有的defer函数执行完毕后，运行时恐慌才会真正的被扩散至调用函数

defer的优势：

1、 对延迟函数的地调用总会在外围函数执行结束前执行。

2、defer语句在外围函数函数体中的位置不限，并且数量不限。

使用defer的注意事项：

1、如果在延迟函数中使用外部变量，就应该通过参数传入。

2、同一个外围函数内多个延迟函数调用的执行顺序，会与其所属的defer语句的执行顺序完全相反。

3、延迟函数调用若有参数传入，那么参数的值会在当前defer语句执行时求出。



小技巧：如果保证不出现死锁的情况，在调用lock之后马上使用defer xx.Unlock()



go语言控制特点：

1、没有do while循环，只有一个更广义的for循环。

2、switch语句灵活多变，还可以用于类型判断。v.(type)

3、if语句和switch语句都可以包含一条初始化子语句（最多一条）。

4、break语句和continue语句可以后跟一条标签（label）语句，以标识需要终止或继续的代码块。

5、defer语句可以使我们更方便的执行异步捕获和资源回收任务。

6、select语句也用于多分支选择，但只与通道配合使用。

7、go语句用于异步启动goroutine并执行指定函数。





