package main

import "fmt"

func main() {
	//统计字符串的长度，按字节进行统计
	var str string = "golang你好" //在golang中，汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str))

	fmt.Println(str)

	//遍历字符串
	for i, val := range str {
		//fmt.Println(i, val) //这样输出的是 acii 码 和 utf8码
		fmt.Printf("%d, %c\n", i, val) // 所以带汉字的字符串尽量使用 Printf 输出
	}
	r := []rune(str) //字符串转切片然后进行遍历
	for i := 0; i < len(r); i++ {
		fmt.Println(i, r[i])
	}

	// new 关键字
	iPtr := new(int)
	*iPtr = 10
	fmt.Println(*iPtr)
	sPtr := new(string)
	*sPtr = "hello"
	fmt.Println(*sPtr)
}
