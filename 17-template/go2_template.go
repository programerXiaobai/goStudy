package main

import "fmt"

/*
通过泛型定义类型：
type Slice[T int | string | float64] []T   使用时需要指定类型：var is Slice[int] = Slice[int]{1, 2, 3}
type MyMap[KEY int | string, VAL any] map[KEY]VAL
type MyStruct[T int | string] struct {
	Id T
	Name string
}
通过泛型定义接口：
type PrintData[T any] interface {}
通过泛型定义通道：
type MyChan[T any] chan T

写泛型的技巧：可以先写 int 类型的函数，然后 int 换成 T，然后加上[]和限定
1. func add(a, b int) int {}
2. func add(a, b T) T {}
3. func add[T int | string](a, b T) T {}
*/

//type SliceInt []int
//type SliceString []string
//type SliceFloat []float64

// 使用泛型 直接等价于上述三条语句。使用的时候需要指定这个泛型是什么类型
type Slice[T int | string | float64] []T

func main() {
	var is Slice[int] = Slice[int]{1, 2, 3}
	fmt.Println(is)

	// 通过泛型定义一个 map
	type MyMap[KEY int | string, VAL any] map[KEY]VAL

	var mp MyMap[string, float64]
	mp = make(MyMap[string, float64])
	mp["zhangsan"] = 2.1
	mp["lisi"] = 1.5
	fmt.Println(mp)
}
