package main

import "fmt"

/*
继承：子类继承父类，只需要把父类的类名写到子类中即可。可以重定义父类的方法，也可以自己添加新方法
go 没有公有继承、私有继承 这些东西。go 只有对外包的 公开 和 隐藏
*/

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()")
}

type SuperMan struct {
	Human //SuperMan 继承 Human，只需要在第一行写上 Human 这个类型即可。go 中是没有 公有继承、私有继承这些东西的
	level int
}

// 重定义父类方法 Eat
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()")
}

// 子类中添加新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()")
}

func (this *SuperMan) Show() {
	fmt.Println(*this)
}

func main() {
	h := Human{"zhangsan", "female"}
	h.Eat()
	h.Walk()

	//定义一个子类对象(需要在里面指定父类 Human 的属性)
	s := SuperMan{Human{"lisi", "male"}, 2}
	s.Eat()
	s.Fly()
	s.Walk()

	//定义子类对象的第二种方法
	var s1 SuperMan
	s1.sex = "male"
	s1.name = "wangwu"
	s1.level = 5
	s1.Fly()

	s1.Show()
}
