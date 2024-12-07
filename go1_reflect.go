package main

import (
	"fmt"
	"reflect"
)

// TypeOf 可以拿到类型信息，有 Name 和 Kind 种类
func reflectType(x interface{}) {
	// 不知道别人调用我这个函数的时候会传进去什么类型的变量
	// 1.通过类型断言，但是只能一个类型一个类型的去断言
	// 2.反射，获取类型
	typeX := reflect.TypeOf(x)
	/*
		在反射中有类型Type 和 种类Kind，因为可以使用 type 关键字去构造很多的自定义类型，而Kind 是底层的类型。有时候我们会需要知道他们的种类，从而提供不同的逻辑
		对于数组、切片、Map、指针，他们的.Name() 都是空
	*/
	fmt.Println(typeX, typeX.Name(), typeX.Kind()) // type.Name获取 Type，type.Kind 获取 Kind。比如对应下面的Person main.Person，Person struct
	fmt.Printf("%T\n", typeX)                      //*reflect.rtype，反射中的一个指针
}

// ValueOf 可以拿到值的信息
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x) // 现在的 v 是 reflect.Value 类型
	fmt.Printf("%v,%T\n", v, v)
	// 将 v 这个 reflect.Value 转换成对应的原类型
	switch v.Kind() { // v.Kind() 可以拿到值对应的类型
	case reflect.Float64:
		// 把 反射取到的值转换成 Float64 的变量
		res := v.Float() // 通过 .Float 方法可以获取 float64的值
		fmt.Printf("float64,%v\n", res)
	case reflect.Float32:
		res := float32(v.Float()) //通过 .Float 方法可以获取 float64的值，强转成 float32
		fmt.Printf("float32,%v\n", res)
	}
}

// 上述方法只能获取值，这个方法可以通过反射拿到值并设置值。因为 go 中是值传递，所以要想修改，需要是指针传递
// 因此调用的时候就是 reflectSetValue(&a)，那么 x 其实就是指针。可以通过 Elem() 方法根据指针取出对应的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 使用 if，或者 switch 也可以
	if v.Elem().Kind() == reflect.Float64 { // 必须加上 v.Elem() 先获取指针的值，再.Kind 获取类型
		v.Elem().SetFloat(2.34)

	} else if v.Elem().Kind() == reflect.Float32 {
		v.Elem().SetFloat(3.14)
	} else if v.Elem().Kind() == reflect.Int8 {
		v.Elem().SetInt(2)
	}
}

type Person struct {
	name string
	age  int
}

func main() {
	var a float64 = 1.23
	var b int8 = 1
	p := Person{"小明", 11}

	//reflectType(a)
	//reflectType(b)
	//reflectType(p)

	//reflectValue(a)
	//reflectValue(b)
	//reflectValue(p)

	reflectSetValue(&a)
	reflectSetValue(&b)
	reflectSetValue(&p)
	fmt.Println(a, b, p)
}
