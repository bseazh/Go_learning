package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

const N = 10
const M = 1000

func read() {
	defer wg.Done()
	rwlock.RLock()
	//lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	//lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 50)
	//lock.Unlock()
	rwlock.Unlock()
}
func main() {

	start := time.Now()
	for i := 0; i < N; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < M; i++ {
		go read()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
