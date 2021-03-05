package main

import (
	"fmt"
	"time"
)

// 程序启动之后会创建一个主 Goroutine去执行;
func main() {
	//单独开启一个goroutine 去执行hello函数（任务）

	//第一种情况会10个里面会出现重复：
	//其原因是因为函数传值时：是"闭包"，外层的for循环累加;取到的i会变化;
	fmt.Println("第一种情况")

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second)

	//第二种情况其实是解决第一种情况所出现的办法
	fmt.Println("第二种情况")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

}
