package main

import "fmt"

//Function A 分两步 还需写成全局的函数
func add(x, y int) int {
	return x + y
}
func Calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func FuncA() {
	fmt.Println("Function : A 分两步")
	x := 1
	y := 2
	fmt.Println("x + y = ", Calc(x, y, add))
}

//Function B 同样是分两步 不过函数写成匿名函数,在函数体内定义
func FuncB() {
	fmt.Println("Function : B  2步到位")
	x := 3
	y := 2
	sub := func(x, y int) int { return x - y }
	Calc2 := func(x, y int, op func(int, int) int) int {
		return op(x, y)
	}
	fmt.Println("x - y = ", Calc2(x, y, sub))
}

//Function C 同样是分两步 在B的基础上,直接执行后一个函数
func FuncC() {
	fmt.Println("Function : C  1.5步到位")
	x := 4
	y := 2
	sub := func(x, y int) int { return x - y }
	res := func(x, y int, op func(int, int) int) int {
		return op(x, y)
	}(x, y, sub)
	fmt.Println("x - y = ", res)

}

//Function D 仅仅是1步完成,两个函数都是匿名并直接返回对应的值
func FuncD() {
	fmt.Println("Function : D  1步到位")
	x := 5
	y := 2
	res := func(x, y int) int {
		return func(x, y int) int {
			return x - y
		}(x, y)
	}(x, y)
	fmt.Println("x - y = ", res)

}
func main() {
	FuncA()
	FuncB()
	FuncC()
	FuncD()
}
