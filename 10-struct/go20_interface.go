package main

import "fmt"

/*
多态：interface。使用 type interface 定义 接口，struct 子类必须实现 父类的所有接口方法，否则无法产生多态
当父类指针(interface 变量) 指向子类对象时，产生多态
*/

// 子类需要继承 interface。interface 本质是一个指针，指向当前 interface 指向的具体类型，还有当前类型包含的函数列表
type AnimalIF interface {
	Sleep()           //睡觉
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

// 具体的类 猫（之前 struct 继承 struct 的时候需要把 父类的类内包含到子类中。但是继承 interface 不需要包含 interface 的名字，只需要把 interface 的方法重写即可）
// 如果子类没有完全重写 interface 的所有方法，则 interface 指针无法指向 这个类，无法实现多态
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 狗
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShouAnimal(animal AnimalIF) {
	animal.Sleep()
}

func main() {
	var animal AnimalIF    //定义一个接口类型，是一个父类指针
	animal = &Cat{"white"} //父类指针指向子类对象
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	animal = &Dog{"yellow"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	ShouAnimal(&Cat{"white"})
	ShouAnimal(&Dog{"yellow"})
}
