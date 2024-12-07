package main

import (
	"fmt"
)

//pair 和 多态 的测试练习

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("ReadBook")
}

func (b *Book) WriteBook() {
	fmt.Println("WriteBook")
}

func main() {
	b := &Book{} //b:  pair<type:Book, value:Book{}地址>

	var r Reader //r:  pair<type, value>
	r = b        //r:  pair<type:Book, value:Book{}地址>
	r.ReadBook()

	var w Writer //w:  pair<type, value>
	w = b        //w:  pair<type:Book, value:Book{}地址>
	w = r.(Writer)
	w.WriteBook()
}
