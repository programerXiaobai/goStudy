package main

import (
	"fmt"
	"sync"
)

/*
多个 goroutine 并发操作全局变量 x，需要加锁
互斥锁：sync.Mutex
Wg sync.WaitGroup
Wg.Add(delta int)   计数器+delta
Wg.Done()           计算器-1
Wg.Wait             阻塞直到计数器变为0

sync.Once：在很多场景下需要一个方法只执行一次，例如只加载一次配置文件，多个 goroutine 关闭通道的时候只关闭一次
sync.Once 只有一个方法 Do：  func (o *Once) Do(f func()) {}   如果执行的函数 f 需要传递参数就需要搭配闭包来使用
*/

var x int
var Wg sync.WaitGroup
var mtx sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		mtx.Lock()
		x++
		mtx.Unlock()
	}
	Wg.Done() // goroutine 运行的函数最后需要加 Done，运行到Done，Wg.Add 中的 delta 会--，有点像信号量
}

func main() {
	Wg.Add(2) //表示开启两个 goroutine
	go add()
	go add()
	Wg.Wait() // 等待上述的两个 goroutine 运行完成
	fmt.Println(x)
}
