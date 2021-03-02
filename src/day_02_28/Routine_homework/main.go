package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)

	for i := 1; i <= 10; i++ {
		go func(i int) {
			<-c1
			fmt.Println("A")
			c2 <- struct{}{}
		}(i)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			<-c2
			fmt.Println("B")
			c1 <- struct{}{}
		}(i)
	}
	c1 <- struct{}{}
	c2 <- struct{}{}

	time.Sleep(3 * time.Second)
}
