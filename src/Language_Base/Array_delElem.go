package main

import (
	"fmt"
	"math/rand"
)

//如何遍历删除 数组里面某个元素
func main() {
	var arr = [7]int{0, 1, 2, 3, 4, 5, 6}

	x := rand.Int() % 7

	fmt.Println("RandomNum = ", x)
	idx := -1
	for i, val := range arr {
		if x == val {
			idx = i
			break
		}
	}
	s1 := arr[:idx]
	s2 := arr[idx+1:]
	s1 = append(s1, s2...)
	//强行转换成数组
	var res = [6]int{}
	for idx, elem := range s1 {
		res[idx] = elem
	}
	fmt.Println("切片: ", s1)
	fmt.Println("数组: ", res)
}
