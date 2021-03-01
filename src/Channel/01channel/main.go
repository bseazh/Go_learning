package main

import "fmt"

var a []int
var b chan int

//channel作为引用类型
//需要Make初始化才能使用;

//其余另外两个是 : slice , map
func main() {
	fmt.Println(b)        // nil
	b = make(chan int, 1) // 不带缓冲区的 通道
	b <- 10
	t := <-b
	println(t)
	b = make(chan int, 16) // 带有缓冲区的 通道

}
