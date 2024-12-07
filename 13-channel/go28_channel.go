package main

import "fmt"

/*
channel：两个 goroutine 之间的通信方式就是 channel
通过 channel 通信的两个 goroutine（称为A，B） 是有同步机制的。A 往 c:=make(chan,int) 中写，B 从 c 中读
无缓冲的 channel：
如果B 读到 c 时还没有数据，说明 A 还没往 c 中写，这时候 B 会阻塞等待，等有数据的时候才会往后执行
如果A 写了 c 后B 没有马上读数据，那么 A 也会阻塞等待，等待 B读完数据之后才会往后执行
所以就是不管 A、B 谁快，到达 channel 都需要等待，等到对方到来才能往后执行
*/

func main() {
	//定义一个 channel。通过 make、chan 来定义，channel 传输的数据类型是 int，无缓存
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行")
		c <- 666 //将666发送给 c
	}()

	num := <-c //从 c 中接收数据，并赋值给 num。也可以 <-c 表示从 c 中读数据，但是并没有使用，而是直接抛弃
	fmt.Println(num)
	fmt.Println("main goroutine 结束")
}
