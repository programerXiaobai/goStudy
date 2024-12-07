package main

import "fmt"

/*
泛型 receiver：泛型带类型的 receiver，将泛型函数绑定到某个类型上
泛型函数：带类型形参的函数
*/

type MySlice[T int | string | float64] []T

// 泛型 receiver：泛型带类型的 receiver，将泛型函数绑定到某个类型上
func (this MySlice[T]) Sum() T {
	var sum T
	for _, v := range this {
		sum += v
	}
	return sum
}

// 泛型函数：带类型形参的函数
func Add[T float64 | int | string](a, b T) T {
	return a + b
}

func main() {
	var ms MySlice[string] = MySlice[string]{"aaa", "bbb", "ccc"}
	strs := ms.Sum()
	fmt.Println(strs)

	fmt.Println(Add[float64](1.2, 3.5))
}
