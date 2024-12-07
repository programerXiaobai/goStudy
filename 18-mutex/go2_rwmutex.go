package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读写锁：sync.RWMutex。当一个 goroutine 获取读锁之后，其他 goroutine 获取读锁会直接进入，写锁会等待；当一个 goroutine 获取写锁之后，其他的 goroutine 无论读锁还是写锁都会等待
*/

var y int64
var wg2 sync.WaitGroup
var mtx2 sync.Mutex    // 花费12.6s
var rwmtx sync.RWMutex //花费10.8s

func Read() {
	//mtx2.Lock()
	rwmtx.RLock()
	time.Sleep(time.Millisecond)
	//mtx2.Unlock()
	rwmtx.RUnlock()
	wg2.Done()
}

func Write() {
	//mtx2.Lock()
	rwmtx.Lock()
	y++
	time.Sleep(time.Millisecond * 10)
	//mtx2.Unlock()
	rwmtx.Unlock()
	wg2.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go Read()
	}

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go Write()
	}

	wg2.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}
