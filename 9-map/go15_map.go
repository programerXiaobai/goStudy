package main

import (
	"fmt"
)

func main() {
	//声明 myMap1 是一种 map 类型，key 是 int，value 是 string。此时只是声明，所以 myMap1是空的
	var myMap1 map[int]string
	if myMap1 == nil {
		fmt.Println("myMap1 是空的")
	} else {
		fmt.Println("myMap1 不是空的")
	}
	//在使用 map 之前，需要给 map 开辟空间，空间的 cap 是10（map 满之后也会动态扩容，扩容也是 2 倍）
	myMap1 = make(map[int]string, 10)
	myMap1[0] = "java"
	myMap1[1] = "C++"
	myMap1[2] = "python"
	fmt.Println(myMap1) //直接打印 map 中的元素

	//第二种方式：使用 := 和 make，此时虽然没有指定容量，但是开辟了空间，是可以直接访问的
	myMap2 := make(map[string]string)
	myMap2["java"] = "one"
	myMap2["C++"] = "two"
	myMap2["python"] = "three"
	fmt.Println(myMap2)

	//第三种方式：直接创建并赋值
	myMap3 := map[string]int{
		"ont":   1,
		"two":   2,
		"three": 3,
	} //主要里面的元素需要加 ,
	fmt.Println(myMap3)
}
