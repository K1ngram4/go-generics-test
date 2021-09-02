package main

// NOTE This won't work since operator == not defined for T
//func findFunc[T any](a []T, v T) int {
//    for i, e := range a {
//        if  e == v {
//            return i
//        }
//    }
//    return -1
//}

//这里泛型是comparable只能比较==，但是any不能比较
func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if  e == v {
			return i
		}
	}
	return -1
}

//func comparableTest[T any](a ,b T) bool {
//	if a == b {
//		return true
//	}
//	return false
//}

func main() {
	print(findFunc([]int{1,2,3,4,5,6}, 5))
	//print(comparableTest("1",1))
}