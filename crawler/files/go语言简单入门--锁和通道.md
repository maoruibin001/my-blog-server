# go语言简单入门--锁和通道 

数据竟态：

数据竟态发生在多个goroutine并发读写一个变量，并且至少一个goroutine在写。

造成竞态条件的根本原因是进程在进行某些操作的时候被中断。虽然在再次运行时会恢复如初，但是外界环境可能在极短时间内发生变化。

锁的分类：互斥锁和读写锁。

使用场景：在大多数goroutine都在读，少部分goroutine在写，此时使用读写锁效率高，否则使用互斥锁效率高。因为读写锁需要更复杂的内部薄记方式。

通道分类：缓冲通道make(chan type, int t)和非缓冲通道make(chan type)。非缓冲通道会阻塞程序，缓冲通道不会。



range在通道循环中有独特的用处，可以在发送通道关闭时自动退出循环。

select：select和switch类似，可以有一个一个的case和一个default。

select和switch区别是：

1、每一个case必须是一个通道。

2、如果没有default且所有case都不满足，则当前goroutine堵塞至有一个case满足为止。

3、如果多个case同时满足条件，select会随机选择一个case执行，而switch会从上到下选择最先满足条件的执行。





