package main

import "fmt"

/*
自定义泛型类型，比如：
type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}
int8 前加~ 表示支持 int8 的衍生类型
其实 any 就是内置定义的一个泛型： type any interface{}
*/

type int8AAA int8 // 给 int8 一个别名

type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}

func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var (
		a int8AAA = 10
		b int8AAA = 15
	)
	fmt.Println(GetMaxNum[int8AAA](a, b)) // 不支持，也就是说int8 的衍生类型是没写到 MyInt 中的，所以不支持，如果想支持，就在 int8 前加~，表示支持 int8 的衍生类型

}
