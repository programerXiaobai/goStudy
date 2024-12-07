package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func func3() {
	fmt.Println("func3")
}

func func_call() {
	defer func1()
	defer func2()
	defer func3()
}

func main() {
	//defer关键字，defer后定义的会在该函数结束之前运行（有点像 C++ 中的析构函数）
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") //defer是压栈的形式，所以后定义的defer先执行

	fmt.Println("hello 1")
	fmt.Println("hello 2")

	func_call()
}
