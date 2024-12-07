package main

import (
	"fmt"
	"reflect"
)

/*
结构体反射很重要，比如配置文件解析，orm 框架等
Field(i int) StructField                           根据索引，返回索引对应的结构体字段信息
NumField() int                                     返回结构体成员字段数量
FieldByName(name string) (StructField, bool)       根据字符串返回对应的结构体字段
FieldByIndex(index []int) StructField              多层成员访问时，根据 []int 提供每个结构体的字段索引，返回字段信息
FieldByNameFunc(match func(string) bool)           根据传入的函数匹配需要的字段
NumMethod() int                                    结构体方法的数量
Method(int) Method                                 返回第 i 个方法
MethodByName(string) (Method, bool)                根据方法名返回方法

获取到 StructField 之后可以继续获取：
type StructField struct {
	Name      string     //字段名字
	PkgPath   string     //包路径
	Type      Type       //字段的类型
	Tag       StructTag  //字段的标签，比如 json
	Offset    uintptr    //字段再结构体中的偏移量
	Index     []int      //用于 Type.FiledByIndex时的索引切片
	Anonymous bool       //是否匿名字段
}
*/

type Student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func (s Student) Study() string {
	msg := "好好学习天天向上"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉快快长大"
	fmt.Println(msg)
	return msg
}

// 根据反射获取结构体方法并调用
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	// 因为下面需要拿到具体的方法去调用，所以使用的是值信息 v，而不是 t
	for i := 0; i < t.NumMethod(); i++ {
		mt := v.Method(i).Type()
		fmt.Printf("Method %v type %v\n", mt.Name(), mt)
		// 通过反射调用方法传递的参数必须是 []reflect.Value{} 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 通过 Call 方法调用结构体的方法
	}

	// 也可以通过方法名获取结构体方法
	t.MethodByName("Sleep") // 返回 Method，bool
	v.MethodByName("Sleep") // 返回 Value
}

func main() {
	stu1 := Student{
		Name:  "小王子",
		Score: 90,
	}

	//通过反射获取结构体中所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// i 就是结构体字段的索引
		f := t.Field(i)
		fmt.Printf("name:%v, type:%v, tag:%v\n", f.Name, f.Type, f.Tag)
		fmt.Println(f.Tag.Get("json"), f.Tag.Get("ini")) // 获取不同的tag
	}

	//根据名字取结构体中的字段
	filed, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v, type:%v, tag:%v\n", filed.Name, filed.Type, filed.Tag)
	}

	printMethod(stu1)
}
