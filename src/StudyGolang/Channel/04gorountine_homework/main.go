package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//3.主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

var wg sync.WaitGroup

func generate_num(ch1 chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func summary_num(ch1 <-chan *job, ch2 chan<- *result) {
	defer wg.Done()
	for {
		newjob := <-ch1
		n := newjob.value
		sum := int64(0)
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: newjob,
			sum: sum,
		}
		ch2 <- newResult
	}
}

func main() {
	//1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan

	wg.Add(1)
	go generate_num(jobChan)

	//2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go summary_num(jobChan, resultChan)
	}

	//3.主goroutine从resultChan取出结果并打印到终端输出
	for x := range resultChan {
		fmt.Printf(" value : %d , sum : %d\n", x.job.value, x.sum)
	}

	wg.Wait()
}
