package main

import "fmt"

/*
泛型：如果用 interface{} 代替泛型，那么函数中会做很多类型断言才能使用
使用泛型： func printArr[T string | int](arr []T)
这里的 []T 就是泛型的形式类型，go 中需要给 T 做一个约束，在函数名后面通过[T string | int] 表示 T 只能是 string 和 int
可以使用 any 表示go 中所有的类型，比如[T any]
可以使用 comparable 表示 go 中所有可以比较的类型，比如 int,uint,float,bool,struct,指针，比如[T comparable]
*/

func main() {
	strs := []string{"wtf", "asd"}
	printArr(strs)
	is := []int{1, 2, 3}
	printArr(is)
}

// 这里的 []T 就是泛型的形式类型，go 中需要给 T 做一个约束，在函数名后面通过[T string | int] 表示 T 只能是 string 和 int
// 可以使用 any 表示go 中所有的类型，比如[T any]
// 可以使用 comparable 表示 go 中所有可以比较的类型，比如 int,uint,float,bool,struct,指针，比如[T comparable]
func printArr[T string | int](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
