package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

//随机数的用法
//1、创建随机种子
//2、rand.Int() -- return int64 的随机数
//3、rand.Intn(x) -- return %x  的随机数

//func f() {
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < 5; i++ {
//		r1 := rand.Int()    // int64
//		r2 := rand.Intn(10) // [0,10)
//		fmt.Println(r1, "  ", r2)
//	}
//}

var Wg sync.WaitGroup

func f1(i int) {
	//任务解决
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println("###:", i)
}
func main() {

	//测试随机数的用法
	//f()

	//利用WaitGroup判断所有的goroutine都结束;
	for i := 0; i < 10; i++ {
		//添加任务
		Wg.Add(1)
		go f1(i)
	}
	//等待任务全部清零 wg 的计数器清零;
	Wg.Wait()
}
