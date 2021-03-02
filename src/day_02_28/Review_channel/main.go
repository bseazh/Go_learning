package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendNum(ch chan int) {
	for {
		x := rand.Intn(10)
		ch <- x
		time.Sleep(time.Second * 3)
	}
}
func main() {

	ch := make(chan int, 1)

	go sendNum(ch)
	for {
		x, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
