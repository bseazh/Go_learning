package main

//参考博客:	https://www.digitalocean.com/community/tutorials/handling-panics-in-go

//log日志登记
import (
	"fmt"
	"log"
)

func main() {

	divideByZero()

	//这句话就是 哪怕发生panic后,还能恢复到并显示出该语句
	fmt.Println("we survived dividing by zero!")

}

func divideByZero() {

	//go中是没有异常捕获的功能,具体到每一个异常都会以err值的形式返回
	// catch 捕获异常并显示异常是什么？
	defer func() {
		// 恢复刚才的发生panic,看看是怎么一回事（捕获异常的类型），同时记录到log日志里
		// err 的值 为 nil 即无异常 , 反之 则有异常出现
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	//try ... 尝试着 1/0
	//执行后goroutine 就会在除以0的位置 panic
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}

//结果展示 :
//2021/01/08 23:02:01 panic occurred: runtime error: integer divide by zero
//we survived dividing by zero!
