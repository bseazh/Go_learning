package main

import "fmt"

//参考博客:https://www.liwenzhou.com/posts/Go/09_function/
//由于defer语句延迟调用的特性;
//defer功能:解决资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

// 个人理解 :
// defer 字面意思 为"延迟",发生时间是在底层
// 1. 返回值的赋值   **2** . defer 语句的执行  3.底层RET的语句执行
//	 备注: ret指令用栈中的数据，修改IP的值，从而实现栈转移(百度百科)
//	 因为defer是配合着匿名函数使用,所以一定程度上是需要回到原函数继续执行.
//	 所以RET是压栈的程序入口，弹出栈顶地址，恢复PC值并并继续执行函数.

// 拿了参考博客的两个面试题进行分析

// defer执行时刻

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
	//1.返回值 = x = 5.
	//2.defer ( x ++ ) x -> 6
	//3.RET
	//4.结果为 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
	//1.(返回值 即 x) = 5.
	//2.defer ( x ++ ) x -> 6
	//3.RET
	//4.结果为 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
	//1.(返回值 即 y) = x = 5.
	//2.defer ( x ++ ) x -> 6
	//3.RET
	//4.结果为 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
	//1.(返回值 即 x) = x = 5.
	//2.defer ( x ++ ) ; x -> 5
	//	原因func是传进去的是 副本 , 也就是值拷贝 , 并没有影响实际 x 的值
	//3.RET
	//4.结果为 5
}
func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
	//1.(返回值 即 x) = x = 5.
	//2.defer ( x ++ ) ; x -> 6
	//3.RET
	//4.结果为 5
}
func Test1() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
}

//Answer:
//5
//6
//5
//5
//6

//结论: 首先看清楚 1.返回值是否为参数
//				2.如果为参数 , defer 改变其值 , 返回值随着改变

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test2() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	// 1. x = 1
	// 2. y = 2
	// 3. defer calc("AA", 1, calc("A", 1, 2))
	// 3.1 .  calc("A", 1, 2) // OutPut: A 1 2 3 ----- ①
	// 3.2 .  defer calc("AA", 1, 3 )
	// 4. x = 10
	// 5. defer calc("BB", x, calc("B", x, y))
	// 5.1 . calc("B", 10, 2) // OutPut: B 10 2 12 -----  ②
	// 5.2 . defer calc("BB", 10 , 12 )

	// 恢复后从后往前弹出执行defer
	// defer calc("BB", 10 , 12 ) //OutPut: BB 10 12 22 -----  ③

	// defer calc("AA", 1, 3 ) //OutPut: AA 1 3 4 -----  ③
}

//Answer:
//A 1 2 3
//B 10 2 12
//BB 10 12 22
//AA 1 3 4

//结论: defer 只能做到最外层的函数延时,里面的函数不能延时
func main() {
	Test1()
	Test2()
}
