package lib1 //属于哪个包就放哪个包

import "fmt"

// 当前 lib1 包 提供的 api
// 函数名的第一个字母如果是大写，说明是对外开放的，其他包是可以调用的；如果是小写，只能在当前包调用
func Lib1Test() {
	fmt.Println("Lib1Test")
}

func lib1Test() {
	fmt.Println("lib1Test")
}

func init() {
	fmt.Println("lib1 init")
}
