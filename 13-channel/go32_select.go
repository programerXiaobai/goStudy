package main

import "fmt"

/*
单流程下一个 go 只能监控一个 channel 的状态，select 可以完成监控多个 channel 的状态。如果 哪个channel 可读或可写，就返回
select {
case <-chan1:
	如果 chan1 成功读到数据，则处理该 case 语句
case chan2 <- 1:
	如果 chan2 成功写入数据，则处理该 case 语句
default:
	如果上面都没成功，则进入 default 的处理
}
上述case监控 chan1 是否可读，chan2 是否可写，谁先触发就处理谁。一般 select 和 for 搭配使用，如果没有 for 循环，触发一次就 select 就结束了
有点像 IO多路复用的 select
*/

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: //判断 c 是否可写，如果可写，就把 x 写进c，同时更新 x、y，这样每次往 c 中写的就是 1，2，3，...斐波那契数列。每次 for 循环都会往 c 中写数据，for 循环直到最后的 quit可读
			x, y = y, x+y
		case <-quit: //quit 可读，说明 子goroutine 结束了，此时结束 for 循环
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //把 c 中的数据打印出来
		}

		quit <- 0 //for 循环结束之后就把 0 给到 quit
	}()

	fibonacii(c, quit) // channel 传形参是传的 引用
}
