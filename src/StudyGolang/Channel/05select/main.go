package main

import "fmt"

//select多路复用
//样例说明:
//1. i = 0 时,当前ch通道为空，A 不能执行，所以只能执行 B
//2. i = 1 时,当前ch通道已满, B 不能执行,所以只能执行 A

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		// case A
		case x := <-ch:
			fmt.Println(x)
		// case B
		case ch <- i:
		}
	}
}
