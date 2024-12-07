package main

import "fmt"

//切片的截取

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:3] //切片，区间是左闭右开，也就是取 s 的 [1,3) 索引的元素，即 [2,3]
	fmt.Printf("slice s1 = %v\n", s1)

	//s[:3]表示从第0个到第3个  s[4:]表示从第4个到末尾
	//s1 和 s 是同一个切片(s1 和 s 是两个指针，这两个指针指向同一个 slice)，改变 s1，s 也会改变
	s1[0] = 100
	fmt.Printf("slice s1 = %v\n", s1)
	fmt.Printf("slice s = %v\n", s)

	//如果害怕改了 s1 会影响 s，那就是使用关键字 copy，可以将底层数组的 slice 一起进行拷贝
	s2 := make([]int, 2)
	copy(s2, s1) //将 s1 拷贝到 s2 中
	fmt.Printf("slice s2 = %v\n", s2)
	s2[0] = 200
	fmt.Printf("slice s2 = %v\n", s2)
	fmt.Printf("slice s = %v\n", s)
	fmt.Printf("slice s1 = %v\n", s1)
}
