package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 导报"testing"，参数输入 *testing.T
//func TestSplit(t *testing.T) {
//	got := Split("a/b/c/d/e/f", "/")
//	want := []string{"a", "b", "c", "d", "e", "f"}
//	// 上面的例子就能通过，下面的就无法通过，因为中午在 utf8 下至少是3个字节，所以 s[idx+1:] 是不对的，应该是 s[idx+len(sep):]
//	//got := Split("我爱你", "爱")   // 函数运行的结果
//	//want := []string{"我", "你"} // 预期的结果
//	// 通过反射判断两个引用类型的变量是否相等
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("got %v want %v", got, want)
//	}
//}

// 单元测试，编写多个测试案例
// 导报"testing"，参数输入 *testing.T
// go test 命令可以查看所有测试案例的情况
// go test -v 可以查看具体测试案例的情况，需要把 t.Run(name, func(t *testing.T){}) 中的匿名函数填入
// 可以使用 go test -run="Split/chinese simple" 表示只跑 chinese simple 这个案例
// go test -cover 可以查看测试通过覆盖率，也就是测试的函数Split 语句运行的代码 占 总代码的比例，越高越好，因为越高说明你的代码没有冗余
// 因为 go test -cover 只输出百分比，如果想详细的看，就先 go test -cover -coverprofile=c.out 表示把覆盖情况按照文件输出，这样该文件夹下就有了 c.out， 然后就查看： go tool cover -html=c.out
func TestSplit(t *testing.T) {
	// 编写测试组
	type test struct {
		s    string
		sep  string
		want []string
	}
	tests := map[string]test{
		"simple1":        test{s: "我爱你", sep: "爱", want: []string{"我", "你"}}, // 测试组中的第一个测试案例，测试程序会把 s,sep 给到 Split 运行出一个切片和 want 比较
		"single sep":     test{s: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"multi sep":      test{s: "abcd", sep: "bc", want: []string{"a", "d"}},
		"chinese simple": test{s: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s falied, Split(%q, %q) = %q, want %q", name, tc.s, tc.sep, got, tc.want)
			}
		})
	}
}

// 性能基准测试
// 比如调用这个函数 1w 次，查看执行的时间，内存占用等
// 需要包含 b *testing.B
// 通过 go test -bench=Split 来测试 Split 函数的性能，重点关注 118.8 ns/op 表示每执行一次消耗118.8ns
// 通过 go test -bench=Split -benchmem 来获取测试的内存信息，其中 112 B/op 表示每执行一次需要112B， 3 allocs/op 表示执行一次需要申请3次内存
func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的树，是按照1，2，5...累加的，也就是先保证Split函数跑一遍，如果不够1s，就再跑，反正肯定要跑够1s
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// 示例代码
// 通过 go test -run Example 运行。记得按照如下的注释写，也就是学院 把 Output 先写下来，每一行注释都写 want 的输出
func ExampleSplit() {
	fmt.Println(Split("沙河有沙又有河", "沙"))
	fmt.Println(Split("a:b:c", ":"))
	// Output:
	// [ 河有 又有河]
	// [a b c]
}
