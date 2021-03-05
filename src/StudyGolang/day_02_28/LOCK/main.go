package main

import (
	"fmt"
	"sync"
)

var x = 0

const N = 5000000

var wg sync.WaitGroup
var lock sync.Mutex //原子锁

func add() {
	for i := 0; i < N; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
