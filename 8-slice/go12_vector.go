package main

import "fmt"

//一般使用过程中很少使用静态数组，都是使用的切片

// 传一个切片数组，方法和静态数组一样的
func printArray2(myArray []int) {
	for _, val := range myArray { //因为 range 返回的是 index 和 val，如果不关心 index，用 _ 表示，表示匿名变量
		fmt.Println("val = ", val)
	}
	myArray[0] = 100 //函数内修改，函数外就也修改了，说明 slice 的传递是引用传递，不是值传递
}

// 动态数组
func main() {
	myArray := []int{1, 2, 3, 4} //[]里面是空的，表示是动态数组，切片 slice
	fmt.Printf("myArray type is : %T\n", myArray)

	printArray2(myArray)

	for _, val := range myArray { //因为 range 返回的是 index 和 val，如果不关心 index，用 _ 表示，表示匿名变量
		fmt.Println("val = ", val)
	}
}
