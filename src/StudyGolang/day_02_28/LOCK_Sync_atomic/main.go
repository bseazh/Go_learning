package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x int64 = 0

func Add() {
	defer wg.Done()
	// 方法一 ：直接加互斥锁
	//lock.Lock()
	//x = x + 1
	//lock.Unlock()

	// 方法二 ： 使用内置的原子操作函数
	atomic.AddInt64(&x, 1)
}
func main() {

	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go Add()
	}
	wg.Wait()
	fmt.Println(x)
}
