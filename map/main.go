package main

import (
	"fmt"
)

//mapFunc 该函数就是将传入的数字每个元素加上<>
func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main() {
	vi := []int{1,2,3,4,5,6}
	vs := mapFunc(vi, func(v int) string {
		return "<" + fmt.Sprint(v) + ">"
	})
	fmt.Println(vs)
}
