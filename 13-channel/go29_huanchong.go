package main

import (
	"fmt"
	"time"
)

/*
有缓冲的 channel：
两个 goroutine A、B，A写B读，因为有缓冲，二者是没有同步的要求的，除非缓冲空了或满了。也就是说
当 channel 满了，再向里面写数据，就会阻塞
当 channel 为空，从里面取数据也会阻塞
*/

func main() {
	c := make(chan int, 3)                              //带有缓冲的 channel，缓冲的长度是3
	fmt.Println("len(c)=", len(c), ", cap(c)=", cap(c)) //打印当前缓冲的 长度 和 容量

	go func() {
		defer fmt.Println("子 goroutine 结束")

		for i := 0; i < 4; i++ { //i有4个，但是channel的len是3，所以加了3个之后，子goroutine 就阻塞了，等到main goroutine 读出一个元素之后 子goroutine会把第四个元素加进去
			c <- i
			fmt.Println("子 goroutine 正在运行，发送的元素=", i, " len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //<- 和 c 中间没有空格
		fmt.Println("num=", num)

	}

	fmt.Println("main goroutine 结束")
}
