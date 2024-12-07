package main

import "fmt"

/*
匿名函数：
在定义匿名函数的时候直接调用，也就是定义之后加上() 进行调用
将匿名函数赋给一个变量（该变量就是函数变量了），通过该变量来调用匿名函数
*/

func main() {
	// 定义匿名函数的时候 调用
	sum := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println("sum = ", sum)

	// 匿名函数赋给变量
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println("sub = ", sub(2, 1))
}
