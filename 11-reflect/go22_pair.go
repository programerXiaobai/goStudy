package main

import "fmt"

/*
go 中，变量是一个 pair 对，即 type 和 value，type 分成 static type 和 concrete type
static type 是 int、string 等， concrete type 是 interface 指向的具体数据类型，系统看得见的类型
*/

func main() {
	var a string
	a = "asf" //现在 a 这个变量 就是一个 pair<static type:string, value:"asf">

	var allType interface{} //此时 allType 也是一个 pair<type, value>，但是 type 指针 和 value 指针并没有指向任务数据
	allType = a             //万能类型指向 a，此时 allType 的type、value指针就指向了数据，即 allType 是 pair<type:string, value:"asf">

	str, _ := allType.(string)
	fmt.Println(str)
}
