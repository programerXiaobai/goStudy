// 四种变量的声明方式
package main

import (
	"fmt"
)

// 声明全局变量，方法一，二，三 都是可以的
// var ga
var gb int = 100
var gc = 200

//方法四声明全局变量(:= 只能用在 函数体内来声明)
//gd := 10

func main() {
	//方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a = ", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型(不是很推荐)
	var c = 10
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c) //Println 是非格式化输出 + 换行， Printf 是格式化输出，%T是输出数据类型
	var cc = "abc"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：（最常用的方法）省去var 关键词，直接自动匹配
	e := 20 //:=是既初始化又赋值
	fmt.Printf("e = %d, type of e = %T\n", e, e)
	ee := "1000"
	fmt.Printf("ee = %s, type of ee = %T\n", ee, ee)
	g := 3.14
	fmt.Printf("g = %f, type of g = %T\n", g, g)

	//fmt.Printf("ga = %d, type of ga = %T\n", ga, ga)
	fmt.Printf("gb = %d, type of gb = %T\n", gb, gb)
	fmt.Printf("gc = %d, type of gc = %T\n", gc, gc)
	//fmt.Println("gd = ", gd)

	//声明多个变量
	var xx, yy int = 100, 200 //同种数据类型就直接写数据类型
	fmt.Printf("xx = %d, yy = %d\n", xx, yy)
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk, ll = 100, "ac" //不同数据类型就不写类型，让自动类型推导
	fmt.Println("kk = ", kk, ", ll = ", ll)
	var ( //不同数据类型通过()写到多行
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
