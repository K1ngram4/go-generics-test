package main

import (
"fmt"
)

//两个参数，一个T类型的数组，一个判断这个T是否是需要的值的函数
func eachFunc[T any](a []T, f func(T) bool) {
	for _, e := range a {
		if !f(e) {
			break
		}
	}
}

func main() {
	// 在123456的有序数组中，找到比4小的数字
	eachFunc([]int{1,2,3,4,5,6}, func(v int) bool {
		fmt.Println(v)
		return v < 4
	})
	// 在字符串数组中找到“”之前的所有字符串
	eachFunc([]string{"foo", "bar", "", "zoo"}, func(v string) bool {
		fmt.Println(v)
		return v != ""
	})
}