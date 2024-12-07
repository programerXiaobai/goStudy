package main

import "fmt"

/*
闭包：一个函数和其相关的引用环境组合的一个整体
比如下面的例子，闭包就是 返回的匿名函数 + 匿名函数意外的变量sum
匿名函数中引用的那个变量会一直保存在内存中，可以一直使用，比如下面例子中的 sum
闭包的本质就是一个匿名函数，只不过是在这个函数引入外界的变量/函数，所以 匿名函数+引用的变量/参数 = 闭包
不建议使用闭包，比如下面的例子可以直接将 sum 拿出作为全局变量，这样一个函数就能解决了，不需要闭包
*/

// 函数功能：求和
// 函数的名字：getSum，参数是空；返回值是一个函数，这个函数的参数是一个 int，返回值也是 int
func getSum() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3
}
