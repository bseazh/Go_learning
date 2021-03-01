package main

import (
	"fmt"
	"strings"
)

//参考博客:https://www.liwenzhou.com/posts/Go/09_function/
// 闭包 = 当前环境 + 参数

// 检验是否以某一个扩展名为后缀,若不是则返回添加扩展名的字符串
// 以.txt为例
//	1. suffix = ".txt" (参数)
//  2. jpgFunc = 函数返回值
//	3. 利用strings.中的一个方法"HasSuffix" 检验是否为后缀
//	4. 如果没有则返回 添加完扩展名后的名称
//	5. fmt( jpgFunf("Test") ) 打印其名称

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestClosure_1() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

// 结论1 ： 此闭包主要凸显以返回值形式接受 用于下一次利用

//------------------------------------------------

// 例2 闭包功能
// Calculate函数 参数为 base 然后返回两个以 base 为基准的 add , sub 函数
// 1. 给出 add 的表达式 参数由 下一次接受 ( 参数 由下一次决定 )
// 1.1 base 在 add 函数得出后 , 顺带修改了当前 base值 ( 当前环境)
// 2. 给出 sub 的表达式 参数由 下一次接受 ( 参数 由下一次决定 )
// 2.1 base 在 sub 函数得出后 , 顺带修改了当前 base值 ( 当前环境)
// 3. 返回 add , sub 两个函数
// 4. base 随之改变 , 每次调用都会改变一次.
// 4.1 base 为什么不是10 , 因为闭包 = 函数 + 外部变量的引用
// 4.2 base 作为Calculate 的参数, 构造 f1 , f2 时 顺带把 base 引用给了它.
// 4.3 所以 base 会随着 f1,f2的执行而改变相应的值

func Calculate(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func TestClosure_2() {
	f1, f2 := Calculate(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
}

//11 = 10 + 1 ;  9 = 11 - 2
//12 =  9 + 3 ;  8 = 12 - 4
//13 =  8 + 5 ;  7 = 13 - 6
func main() {
	TestClosure_1()
	TestClosure_2()
}
