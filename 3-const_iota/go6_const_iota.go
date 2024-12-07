package main

import "fmt"

// const 可以定义枚举类型
const (
	//可以在const() 添加一个关键字iota，每行的 iota 都会累加1，第一行的 iota 的默认值是 0
	BEIJING  = iota //iota = 0
	SHANGHAI        //iota = 1
	SHENZHEN        //iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2  所以 a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2  所以 c = 2, d = 3
	e, f                      // iota = 2

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3   所以 g = 6, h = 9
	i, k                      // iota = 4
)

func main() {
	//常量（只读）
	const length int = 10
	//var a int = 20
	fmt.Println("length:", length)
	//fmt.Println("a:", a)

	fmt.Println("BEIJING:", BEIJING)
	fmt.Println("SHANGHAI:", SHANGHAI)
	fmt.Println("SHENZHEN:", SHENZHEN)

	fmt.Println("a:", a)

}
