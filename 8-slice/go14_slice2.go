package main

//切片的扩容机制

import "fmt"

func main() {
	//表示定义的切片 len 是3，容量 capability 是5
	var nums = make([]int, 3, 5)
	fmt.Printf("nums len:%d, cap:%d, slice:%v\n", len(nums), cap(nums), nums)

	//向切片中 追加元素1，此时 nums 的 len 是4，有点像 C++ 中 vector 的push_back
	nums = append(nums, 1)
	fmt.Printf("nums len:%d, cap:%d, slice:%v\n", len(nums), cap(nums), nums)

	nums = append(nums, 2)
	fmt.Printf("nums len:%d, cap:%d, slice:%v\n", len(nums), cap(nums), nums)

	//如果 slice 的 len == cap 了，此时再追加，会动态扩容，是扩容两倍的 cap，类似于 vector
	nums = append(nums, 3)
	fmt.Printf("nums len:%d, cap:%d, slice:%v\n", len(nums), cap(nums), nums)

	fmt.Println("--------------------")

	//创建一个 slice 的时候没有指定容量，则 cap 就是 len
	nums2 := make([]int, 3)
	fmt.Printf("nums2 len:%d, cap:%d, slice:%v\n", len(nums2), cap(nums2), nums2)
	//此时给 nums 追加元素，仍然是两倍的扩容
	nums2 = append(nums2, 1)
	fmt.Printf("nums2 len:%d, cap:%d, slice:%v\n", len(nums2), cap(nums2), nums2)
}
