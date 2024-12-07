package main

import "fmt"

// type 和 struc 定义一个结构体
type Book struct {
	id     int
	title  string
	author string
}

// struct 传参，但是是值传递
func printBook(book Book) {
	book.id = 2
	fmt.Println(book)
}

// struct 指针传递
func changeBook(book *Book) {
	//标准来说应该是加解引用的，但是省略* 也是可以的。因此虽然是指针传递，其实使用的时候可以按照引用传递来使用，也就是模板参数中加*，函数体内就什么都不用加了
	(*book).author = "lisi"
	book.id = 3
}

func main() {
	var book1 Book
	book1.id = 1
	book1.title = "golang"
	book1.author = "zhangsan"

	fmt.Printf("%v\n", book1)

	printBook(book1)

	changeBook(&book1)
	fmt.Printf("%v\n", book1)
}
