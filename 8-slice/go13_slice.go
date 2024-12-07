package main

import "fmt"

//声明切片的四种方式

func main() {
	//声明 slice1 是一个切片，默认值是1，2，3.  长度 len 是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1) //%v表示打印所有信息

	//声明 slice2 是一个切片，但是并没有给 slice 分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	slice2 = make([]int, 3) //使用 make 给 slice2 开辟空间，之后才可以访问 slice2[0]
	slice2[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//声明 slice3 是一个切片，同时给 slice 分配空间，初始化值是 0
	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	//和方法三类型，使用 := 推导（这种方式用的最多）
	slice4 := make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	//判断 slice 是否为空
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
