/*
	课堂作业：
		既然导入包时可以j进行简写，那么声明多个变量、常量、全局变量
	或者一般类型（非接口、非结构）是否也可以用同样的方法呢？
*/
// 当前程序的包名
package main

// 导入其它的包
import "fmt"

// 常量的定义
const (
	PI = 3.14
	E  = 2.714
)

// 全局变量的声明与赋值
var (
	n0 = "0"
	n1 = "1"
	n2 = "2"
)

// 一般类型声明
type (
	newType  int
	newType2 float32
	newType3 string
)

// 由 main 函数作为程序入口点启动
func main() {
	//局部变量
	var (
		a int = 1
		b int = 2
		c int = 3
	)
	fmt.Println(a, b, c)
	fmt.Println("Hello world!你好，世界！")
}
