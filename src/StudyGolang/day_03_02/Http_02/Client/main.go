package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func test1() {
	resp, err := http.Get("http://127.0.0.1:9090/Path2/?name=hhz&age=18")
	if err != nil {
		fmt.Printf("Http get failed , err : %v\n", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body) // 在客户端中接收从服务端发来的Body消息
	if err != nil {
		fmt.Printf("Read resp.Body failed , err : %v\n", err)
		return
	}
	fmt.Println(string(b))
}

// false : 共用一个clients适用于 请求比较频繁
// true : 定义在局部变量即可 适用于 请求不是特别频繁，用完就关闭
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func test2() {
	//Parse解析 url 返回其url对象
	urlObj, _ := url.Parse("http://127.0.0.1:9090/Path2/")

	//构建一个请求对象
	data := url.Values{}
	data.Set("name", "新垣结衣")
	data.Set("age", "18")

	//编码 以确保在不同的服务器解析的编码相统一
	queryStr := data.Encode()
	fmt.Println(queryStr)

	//将请求对象的访问字段 进行赋值
	urlObj.RawQuery = queryStr

	//利用请求对象 加上 方法 返回其新的request(询问)
	req, err := http.NewRequest("GET", urlObj.String(), nil)

	//自定义其Client发送req, 接收其从服务端返回的 response(响应)
	//resp, err := client.Do(req) 使用全局的client
	resp, err := http.DefaultClient.Do(req)

	b, err := ioutil.ReadAll(resp.Body) // 在客户端中接收从服务端发来的Body消息
	if err != nil {
		fmt.Printf("Read resp.Body failed , err : %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(string(b))
}

func main() {
	//test1()
	test2()

}
