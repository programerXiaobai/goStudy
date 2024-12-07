package main

import (
	"fmt"
	"sync"
)

/*
因为内置的 map 并不是并发安全的，如果并发的写或读就需要加上读写锁
sync.Map:开箱即用的并发安全版 map。不用像 map一样使用 make 函数初始化就能直接使用。同时 sync.Map 内置了 Store，Load，LoadOrStore，Delete，Range 等方法
Store：原子存，类似于 set
Load：原子取，类似于 get
LoadOrStore：先获取，如果不存在，就保存
Delete：删除
Range：遍历
sync.Map 是 interface{} 类型，所以是不需要指定类型的，直接用
*/

var m = make(map[int]int)
var wg3 sync.WaitGroup
var m2 sync.Map

func get(key int) int {
	return m[key]
}

func set(key, val int) {
	m[key] = val
}

func main() {
	//for i := 0; i < 30; i++ {
	//	wg3.Add(1)
	//	go func(i int) {
	//		defer wg3.Done()
	//		set(i, i+100)
	//		fmt.Printf("key:%v, value:%v\n", i, get(i))
	//	}(i)
	//}
	//wg3.Wait()

	for i := 0; i < 30; i++ {
		go func(i int) {
			wg3.Add(1)
			m2.Store(i, i+100)
			val, _ := m2.Load(i)
			fmt.Printf("key:%v, value:%v\n", i, val)
			wg3.Done()
		}(i)
	}
	wg3.Wait()
}
