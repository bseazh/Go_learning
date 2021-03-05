package main

import (
	"flag"
	"fmt"
	"time"
)

// flag args

func f1() {
	// 创建一个标志位参数
	name := flag.String("name", "hhz", "请输入姓名")
	age := flag.Int("age", 18, "请输入年龄")
	married := flag.Bool("married", false, "婚否？")
	cTime := flag.Duration("ct", time.Hour, "结婚多长时间")

	//使用 flag 必须要先 解析
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
}

//go run main.go
//hhz
//18
//false
//1h0m0s

//go run main.go -name 靓仔 -age=25 -married=true -ct 1000h
//靓仔
//25
//true
//1000h0m0s
func f2() {
	var (
		name    string
		age     int
		married bool
		cTime   time.Duration
	)
	flag.StringVar(&name, "name", "hhz", "请输入姓名")
	flag.IntVar(&age, "age", 18, "请输入年龄")
	flag.BoolVar(&married, "married", false, "婚否？")
	flag.DurationVar(&cTime, "ct", time.Hour, "结婚多长时间")
}

func f3() {
	// 创建一个标志位参数
	name := flag.String("name", "hhz", "请输入姓名")
	age := flag.Int("age", 18, "请输入年龄")
	married := flag.Bool("married", false, "婚否？")
	cTime := flag.Duration("ct", time.Hour, "结婚多长时间")

	//使用 flag 必须要先 解析
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数 , 以[]string类型
	fmt.Println(flag.NArg())  //返回命名行参数后的其他参数的个数
	fmt.Println(flag.NFlag()) //返回使用命令行参数的个数
}

//go run main.go -name 靓仔 -age=200 a b c d
//靓仔
//200
//false
//1h0m0s
//[a b c d]
//4
//2

// flag 获取命令行参数
func main() {
	//f1()
	//f2()
	f3()
}
