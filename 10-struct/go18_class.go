package main

/*
封装：使用 struct 定义 类，struct 内只能写成员变量，成员函数需要写到外部，通过 (this* structName) 来表示这个函数是成员函数
根据 成员函数、成员变量 的首字母是大写还是小写 表示 是 public 还是 private
*/

import "fmt"

// struct 的名字 Hero 的首字母大写，表示其他包可以访问，小写的话其他包是无法访问的
type Hero struct {
	//如果类的属性首字母大写，表示该属性其他包是可以访问的，否则只能在该包访问（所以 go 是以首字母的大小写来表示 public 和 private 的）
	name  string
	Age   int
	level int
}

// (this Hero) 表示当前方法绑定到 Hero 结构体中，也就是说 这个方法是 Hero 的成员方法
func (this Hero) GetName() string {
	defer fmt.Println("name = ", this.name)
	return this.name //可以通过 this 调用 Hero 这个结构体的成员变量
}

func (this Hero) SetName(newName string) {
	//setName 其实没有修改了 name。因为 this 是调用该方法对象的一个拷贝，即 hero 的拷贝
	this.name = newName
}

// 所以一般定义成员函数，都是需要是 (this* Hero) 的
func (this *Hero) SetAge(newAge int) {
	// 使用指针，这样就能真正修改了
	this.Age = newAge
}

func main() {
	//创建一个对象
	hero := Hero{name: "zhangsan", Age: 18, level: 1}
	name := hero.GetName()
	fmt.Println(name)
	hero.SetName("lisi") //setName 其实没有修改了 name
	fmt.Println(hero)

	hero.SetAge(20)
	fmt.Println(hero)

	fmt.Println(hero.name)
	fmt.Println(hero.Age)
}
