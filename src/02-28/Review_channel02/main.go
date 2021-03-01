package main

import (
	"fmt"
	"time"
)

var N = int(8)
var M = int(3)
var notifyChan = make(chan struct{}, N)

func Salve_task(id int, tasks <-chan int, results chan<- int) {
	for j := range tasks {
		time.Sleep(time.Second)
		fmt.Printf("worker(%d) solve task(%d)\n", id, j)
		results <- j * 2
		// 创建信号
		notifyChan <- struct{}{}
	}
}

func main() {
	tasks := make(chan int, N)
	results := make(chan int, N)

	// 创建 N 个任务
	go func() {
		for i := 0; i < N; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// 结束任务后把信号回收
	go func() {
		for i := 0; i < N; i++ {
			<-notifyChan
		}
		close(results)
	}()

	// 开启 M 个goroutine 执行任务
	go func() {
		for i := 1; i <= M; i++ {
			go Salve_task(i, tasks, results)
		}
	}()

	//输出结果
	for x := range results {
		fmt.Println(x)
	}
}
