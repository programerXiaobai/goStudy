// init 函数比 main 函数优先运行，一般是用来初始化变量的
package main

import (
	//_ "code/src/goStudy/5-init/lib1" //""前加_ 表示导入匿名包，因为导报 lib1 之后就必须使用这个包，但是匿名包就不能用包里面的方法，但是会执行包中的 init 方法
	"code/src/goStudy/5-init/lib1"
	"code/src/goStudy/5-init/lib2"
	mylib2 "code/src/goStudy/5-init/lib2" //这是给 lib2 这个包起别名，别名是 mylib2
	// . "code/src/goStudy/5-init/lib2" //""前加. 表示把 lib2 这个包中的东西全加载到这个包中了，也就是调用Lib2Test 的时候不用加 lib.  直接调用方法就行。但是不要轻易使用，容易引起方法重名
)

func main() {
	lib1.Lib1Test() //调用的时候就是使用 lib1 包里面的 Lib1Test 方法
	lib2.Lib2Test()
	mylib2.Lib2Test()

	//lib1.lib1Test() //调用不了其他包中小写开头的函数
}
