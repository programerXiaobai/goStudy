package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
通过 go + 匿名函数 创建 goroutine   需要注意的是 匿名函数之后需要加() 表示调用
通过 go + 有参匿名函数 创建 goroutine   需要注意后面的()里面需要指定参数
通过 runtime.Goexit() 退出当前协程
*/

func main() {
	//用过 go + 一个形参为空，返回值为空的匿名函数 创建 goroutine
	go func() {
		defer fmt.Println("A.defer")

		//return //在 goroutine 还没运行完成之前就退出

		func() {
			defer fmt.Println("B.defer")

			//return //只是退出 这个匿名函数，并没有退出整个 goroutine
			//如果需要退出当前的整个 goroutine，需要使用以下方法
			runtime.Goexit()

			fmt.Println("B")
		}() //如果没有后面的()，那么只是匿名函数的声明，并没有被调用，加上后面的()就是被调用了。如果这个匿名函数是有参数的，那么最后的()里面也需要放参数

		fmt.Println("A")
	}() //同样的，匿名函数后面加() 表示调用

	//通过有参的匿名函数创建 goroutine。如果想在 main goroutine 中获取 子 goroutine 中的返回值，需要使用 channel
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ",b = ", b)
		return true
	}(10, 20)

	//main goroutine 中设置一个死循环使得 不退出
	for {
		time.Sleep(time.Second)
	}
}
