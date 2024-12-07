package main

import "fmt"

// go 默认是值传递，要想使用类似于C++的指针传递或引用传递，可以使用指针（一般情况下不使用指针，只有这种情况才会使用）
func changeVal(p int) {
	p = 10
}

// 指针传递
func changePointer(p *int) {
	*p = 10
}

// 交换两个数据
func swap(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	a := 1
	changeVal(a)
	fmt.Println("a = ", a)

	changePointer(&a)
	fmt.Println("a = ", a)

	var b int = 10
	var c int = 20
	swap(&b, &c)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	//二级指针。一般很少用
	var p *int
	p = &a
	fmt.Println("p = ", p)
	fmt.Println("&a = ", &a)
	var pp **int //二级指针
	pp = &p
	fmt.Println("pp = ", pp)
}
