package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("B: %d\n", i)
	}
}
func main() {
	//默认CPU的逻辑核心数,默认跑满整个CPU
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
