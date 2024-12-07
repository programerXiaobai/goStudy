package main

import "fmt"

// map 传参，此时 m 是引用传递，m 和 cityMap 指向同一块内存空间
func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func changeVal(cityMap map[string]string) {
	cityMap["uk"] = "london"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "newyork"

	//遍历，使用 range，返回的是 key，val
	for key, val := range cityMap {
		fmt.Println("key:", key, "val:", val)
	}
	fmt.Println(cityMap)

	//删除，传入key
	delete(cityMap, "china")

	//修改
	cityMap["usa"] = "dc"

	for key, val := range cityMap {
		fmt.Println("key:", key, "val:", val)
	}

	changeVal(cityMap)
	printMap(cityMap)
}
