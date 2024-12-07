package main

/*
通过 go 关键字创建一个 goroutine，运行 goroutine 的时候不会阻塞 main goroutine
但是 main 线程退出之后所有的 goroutine 就都退出了
下面展示 通过 go + 函数名字 创建一个 goroutine
*/

import (
	"fmt"
	"sync"
	"time"
)

// 创建一个任务方法
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 通过 sync.WaitGroup 使得主协程等到从协程运行完成
func worker(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
}

func main() {
	//newTask() //直接允许任务就是主线程允许这个任务，主 goroutine 的后面操作会被阻塞

	//通过 go 关键字创建一个 goroutine，执行 newTask() 流程。主 goroutine 后面的操作就不会被阻塞了
	go newTask()

	// 如果没有下面的操作，主 goroutine 创建完从 goroutine 的 newTask 之后就退出了，因此 从 goroutine 也不会运行
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go worker(wg)
}
