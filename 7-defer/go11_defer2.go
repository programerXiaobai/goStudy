package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

// 如果 defer 和 return 在一个函数中，是 return 先执行 defer 后执行（defer 是函数生命周期结束之后才会执行）
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
