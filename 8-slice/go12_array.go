package main

import "fmt"

// 数组传参，必须表面是[10]int，因此如果是[5]int，需要再写一个函数
// 这样写仍然是一个值拷贝过程
func printArray(a [10]int) {
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func main() {
	//定义固定长度的数组
	var myArray [10]int

	for i := 0; i < 10; i++ {
		fmt.Println(myArray[i])
	}
	//遍历的时候可以使用 len 获取长度
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	myArray2 := [10]int{1, 2, 3, 4} //前四个值是1，2，3，4， 后面的是初始值0
	//遍历的时候可以使用 range，返回的是两个值，分别是 索引index 和 值value
	for index, value := range myArray2 {
		fmt.Println(index, value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray types = %T\n", myArray)
	fmt.Printf("myArray2 types = %T\n", myArray2)

	printArray(myArray2)
}
