package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m1 = make(map[string]int)
var lock sync.Mutex
var wg sync.WaitGroup

func get(key string) int {
	return m1[key]
}

func set(key string, value int) {
	m1[key] = value
}

func Test1() {
	//fatal error: concurrent map writes
	//出现其现象是因为 map本身不支持并发安全
	//解决办法1：直接加互斥锁
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(x int) {

			key := strconv.Itoa(x)
			lock.Lock()
			set(key, x)
			lock.Unlock()

			lock.Lock()
			fmt.Printf("key:%v,value:%v\n", key, get(key))
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//解决办法2: 直接使用sync中所提供的sync.Map中内置函数
var m2 = sync.Map{}

func Test2() {
	//fatal error: concurrent map writes
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(x int) {
			key := strconv.Itoa(x)

			//必须使用 sync.Map内置Store方法设置键值对
			m2.Store(key, x) //set(key, x)

			//必须使用 sync.Map内置的Load方法进行查看
			value, _ := m2.Load(key)
			fmt.Printf("key:%v,value:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main() {
	//Test1()
	Test2()
}
