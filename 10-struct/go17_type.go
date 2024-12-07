package main

import "fmt"

// type 用来给数据类型弄别名，比如 int64 定义成 ll
type ll int64

func main() {
	var a ll = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
