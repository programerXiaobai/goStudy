package main

import (
	"fmt"
	"reflect"
)

/*
反射 就是 根据变量a 获取他的 type 和 value。这样可以获取未知变量的 type 和 value
源码： func ValueOf(i interface{}) Value {...}  //输入任意数据类型，获取数据 value，如果传入的 interface 是空的，返回就是 0
      func TypeOf(i interface{}) Type {...}    //输入任意数据类型，获取类型 type，如果传入的 interface 是空的，返回就是 nil
*/

func reflectNum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", this)
}

// 获取复杂数据类型的 type 和 value
func DoFiledAndMethod(input interface{}) {
	//获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is ", inputType)

	//获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue)

	//通过 type 获取里面的字段
	//1.先通过 reflect.TypeOf 获取 type，通过这个 type 获取 NumField，进行遍历
	//2.得到每个 field 即 struct 中的每个字段的数据类型之后，每个 field 的 Interface{} 方法获取对应的 value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)              //通过 Field 获取第 i 个字段
		value := inputValue.Field(i).Interface() //通过 Interface{} 获取对应的 value

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通过 type 获取里面的 方法 并 调用
	// 通过 reflect.Type 获取 type 之后，通过 NumMethod，进行遍历
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "Alex", 18}
	DoFiledAndMethod(user)
	//user.Call()
}
