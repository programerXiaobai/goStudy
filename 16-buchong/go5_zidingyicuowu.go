package main

import (
	"errors"
	"fmt"
)

/*
异常处理：自定义错误
使用 errors 包中 New 来自定义一个错误   func New(str string) error
比如 err := errors.New("除数不能为0")，这个函数可以返回 err error 即可，其他函数可以接收这个返回值
这样其他函数捕获异常之后，打印这个异常，后面的逻辑仍然能够执行
但是如果这个异常很重要，不想执行后面的逻辑的话，就使用 builtin 包中的 panic 函数，比如 panic(err) 直接退出程序
*/

// 返回一个 error 类型
func test2() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误
		err := errors.New("除数不能为0")
		return err
	} else {
		res := num1 / num2
		fmt.Println(res)
	}
	return nil
}

func main() {
	err := test2()
	if err != nil {
		fmt.Println(err)
		panic(err) // 直接退出程序
	}
	fmt.Println("后面的逻辑...")
}
