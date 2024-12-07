package main

import "fmt"

/*
go 中方法作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型struct 都可以有方法，int,float 等也可以有
*/

type integer int

func (i integer) print() {
	fmt.Println(i)
}

func main() {
	var i integer
	i = 20
	i.print()
}
