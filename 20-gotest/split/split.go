package main

import "strings"

// 字符串切割

// 将 s 按照 sep 进行分割，返回一个字符串的切片
// 比如 Split("我爱你", "爱") 返回 []string{"我", "你"}
func Split(s, sep string) (res []string) {
	res = make([]string, 0, strings.Count(s, sep)+1) // 因为基准测试的时候发现每执行一次需要申请三次内存，说明 append 的时候经常需要动态扩容，直接申请对应的 cap 就可以避免动态扩容
	idx := strings.Index(s, sep)
	for idx > -1 {
		res = append(res, s[:idx])
		//s = s[idx+1:]
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	res = append(res, s)
	return
}
