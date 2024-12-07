package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
加锁会涉及到内核态的上下文切换，会比较耗时，代价高
针对基本数据类型，可以使用原子操作来保证并发安全，go 中的原子操作在用户态就可以完成，因此性能更好
sync/atomic 的4个操作
AddXXXType：原子的增减，比如 AddInt64,AddInt,AddUint,AddPointer
LoadXXXType: 取值
StoreXXXType：存值
CAS，CompareAndSwapInt(addr *int, old, new int) bool: 比较 addr 指向的值与 old 是否相等，相等则将 addr值 改为 new并返回 true；否则什么都不做 返回 false
			当有大量的 goroutine 对变量进行读写操作时，可能导致 CAS 操作无法成功，这是可以利用 for 循环多次尝试
*/

func main() {
	var count int64
	var wg4 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg4.Add(1)
		go func() {
			atomic.AddInt64(&count, 10)
			wg4.Done()
		}()
	}
	wg4.Wait()

	fmt.Println(count)
}
