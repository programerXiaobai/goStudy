package main //当前程序所在的包。只要当前这个文件包含main函数的，就需要包含main包

import (
	"fmt"
	"time"
) //导入标准输入输出包 和 时间包

func main() {
	//加不加;都可以，建议不加
	fmt.Println("hello world")
	time.Sleep(1 * time.Second) //睡眠一秒
}
