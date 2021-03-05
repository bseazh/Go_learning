package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 生成100个数 放进 ch1
// 2. 从 ch1 中取数 平方后放进 ch2
// 3. 从 ch2 中打印

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- i * i
	}
	once.Do(func() { close(ch2) })
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(3)

	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)

	// range 的使用范围特殊在于 ch2 关闭即可 跳出循环
	// 同时 range 遍历时会把通道中的数字取出
	for ret := range ch2 {
		fmt.Println(ret)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		t := <-ch2
		fmt.Println(t)
	}
}
