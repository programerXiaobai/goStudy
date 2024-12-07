package lib2

import "fmt"

// 当前 lib2 包 提供的 api
func Lib2Test() {
	fmt.Println("Lib2Test")
}

func init() {
	fmt.Println("lib2 init")
}
