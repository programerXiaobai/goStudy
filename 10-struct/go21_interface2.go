package main

import "fmt"

/*
interface 除了可以作为接口实现多态，还可以实现作为 空接口、通用万能类型
int,string,float32,float64,struct 等都实现了 interface{}。也就是可以用 interface{} 类型 引用 任意的数据类型
可以通过 arg.(string) 这样的断言机制判断 interface{} 是否是字符串，返回 val 和 ok，ok 表示是不是，如果是的话， val 是字符串的值
interface{} 本质是值传递，如果想在函数中修改值，函数外的实参也修改的话，就断言的时候取出指针，相应的，传递到这个函数的时候传的是地址
*/

// interface{} 是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println("arg is ", arg)

	//interface{} 如何区分 此时引用的底层数据类型是什么
	//go 语言 给 interface{} 提供了 “类型断言” 机制，用来判断 interface 是哪个类型的
	value, ok := arg.(string) //通过.(string) 判断 agr 是不是 string，返回 value 和 ok，ok 表示是不是，如果是的话， value 是字符串的值
	if !ok {
		fmt.Println(" arg is not a string")
	} else {
		fmt.Println("arg is string and value = ", value)
	}
}

type Book1 struct {
	author string
}

// interface{} 本质是值传递，如果想在函数中修改值，函数外的实参也修改的话，就断言的时候取出指针，相应的，传递到这个函数的时候传的是地址
func change2(a interface{}) {
	if val, ok := a.(*int); ok {
		*val = 10
	}
}

func main() {
	book := Book1{"golang"}
	myFunc(book)

	myFunc(100)

	myFunc("www")

	a := 0
	change2(&a)
	fmt.Println(a)
}
