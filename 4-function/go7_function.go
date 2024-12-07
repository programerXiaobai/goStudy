package main

import "fmt"

// 参数模板中的参数的类型在参数的后面，返回值在参数模板的后面
func foo1(a int, b string) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 函数返回多个值，匿名的
func foo2(a int, b string) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

// 返回多个返回值，有形参名称的
func foo3(a int, b string) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值遍历赋值
	r1 = 100
	r2 = 200
	return //直接返回就行
}

// 如果返回值都是一样的，只后面写类型也可以
func foo4(a int, b string) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1,r2 属于foo4的形参，初始化默认值是0
	//r1,r2 作用域空间 是foo4， 整个函数体{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回值遍历赋值
	r1 = 100
	r2 = 200
	return //直接返回就行
}

func main() {
	c := foo1(555, "abc")
	fmt.Println("c = ", c)

	res1, res2 := foo2(565, "abddc")
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)

	res1, res2 = foo3(333, "foo3")
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)

	res1, res2 = foo4(444, "foo4")
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)
}
