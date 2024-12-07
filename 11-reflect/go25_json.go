package main

import (
	"encoding/json" //导入json包
	"fmt"
)

/*
结构体标签在 json 中的应用:
在结构体成员的后面加上标签，标签的 key 是 json，val 是 这个结构体的成员在 json 中 显示的字段名字（注意结构体的成员变量要大写，否则这个成员变量不在 json 中）
导入 json 包：import "encoding/json"
使用 json.Marshal 进行编码，把结构体变量 编码成 json，返回 jsonStr 和 err 两个返回值
使用 json.Unmarshal 进行解码，把 json 解码成 结构体变量，返回 err
*/

// 如果需要把结构体转josn，需要在每个变量后面加标签`json:tag`，这个 tag 就是这个结构体成员在json中显示的名字
type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"rmb"`
	Actor []string `json:"actor"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "张柏芝"}}

	//编码： 结构体 ----> json
	jsonStr, err := json.Marshal(movie) //通过 json 库中的 Marshal 将 movie 转 json，返回两个变量：json 字符串和 error code. Marshal 会把结构体成员的标签作为json的字段
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("%s\n", jsonStr)

	//解码：  json ----->结构体
	//jsonStr := {"title":"喜剧之王","year":2000,"rmb":10,"actor":["zhouxingchi","张柏芝"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie) //通过 json 库的 Unmarshal 进行解码，返回一个变量err
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
