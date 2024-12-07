package main

import (
	"fmt"
	"reflect"
)

// 允许给结构体的成员变量绑定一些标签，比如 name 绑定了两个标签，sex绑定了一个标签，标签用 反引号 括起来
/*
使用的时候可以通过变量获取到标签，作用：其他包导入的这个包的时候，如果使用这些属性，那么会看到标签，标签就是对这个属性的说明，功能等
可以应用于 json 转换，类似于注解
*/
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(s interface{}) {
	t := reflect.TypeOf(s).Elem() //获取当前 struct 的所有元素

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info") //获取当前第 i 行的 Tag 的，get 到 info 这个 tag
		fmt.Println("info: ", tagInfo)

		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("doc: ", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
