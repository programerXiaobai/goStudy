package main

import "fmt"

/*
错误处理
如果程序出错，会报 panic，后面跟的就是错误，比如 runtime error：integer divide by zero
错误处理的捕获机制：defer + recover，一般使用defer+匿名函数，匿名函数中写 recover 捕获异常
异常捕获之后后面的代码仍然会执行
*/

func test() {
	//利用 defer + 匿名函数 + recover 来捕获异常
	defer func() {
		//调用 recover 内置函数，可以捕获异常
		err := recover()
		// 如果没有捕获异常，返回值是 nil；有异常，则返回异常
		if err != nil {
			fmt.Println("异常已捕获，err:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	test()
	fmt.Println("后面的逻辑...")
}
