package main

import "fmt"

/*
关闭 channel：使用 close 函数关闭 channel。关闭 channel 时需要注意此时生产者不需要往 channel 中放数据的时候才能关闭
可以消费者可以通过 data,ok := <-channel 获取 data 和 ok，如果 channel关闭了，那么ok就是 false，此时 消费者就也可以结束读取 channel 的逻辑了
如果关闭 channel，无法向 channel 发数据，否则会出现错误：panic: send on closed channel
关闭 channel 只是表示不会向 channel中放数据，消费者仍然可以读数据，读完之后 channel 才会关闭

channel 与 range：通过 range 来迭代不断操作 channel
for data := range c
从c 中读数据给 data。如果 c中有数据，就返回；如果c 中没有数据，range就阻塞；如果 c关闭，range 就退出
*/

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//通过 close 关键字关闭一个 channel
		close(c)
	}()

	for {
		//if语句的缩写，先执行; 前面的语句“data, ok := <-c”，然后判断 ok 是否为 true
		if data, ok := <-c; ok { //从channel中回去的 data 表示管道的数据，ok 表示是否获取成功
			fmt.Println(data)
		} else { //从管道中读取数据失败，说明管道已经被关闭了，则直接 break for 循环
			break
		}
	}

	//上述 for 循环代码可以 通过 range 实现，即通过 range 来迭代不断操作 channel
	//从c 中读数据给 data。如果 c中有数据，就返回；如果c 中没有数据，range就阻塞；如果 c关闭，range 就退出
	//for data := range c {
	//	fmt.Println(data)
	//}

	fmt.Println("Finished...")
}
