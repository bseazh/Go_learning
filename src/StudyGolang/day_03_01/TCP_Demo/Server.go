package main

import (
	"fmt"
	"net"
)

// TCP server端
//1.监听端口
//2.接收客户端请求建立链接
//3.创建goroutine处理链接。

// 处理函数
func process(conn net.Conn) {

	defer conn.Close() // 关闭连接
	// 3 . 与客户端通信
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from conn failed , err:", err)
			return
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

// Server 端
func main() {
	// 1 . 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Start Tcp server on 127.0.0.1:9000, err:", err)
		return
	}
	fmt.Println("监听中:127.0.0.1:9000")
	// 2 . 等待客户端来跟我连接
	for {
		conn, err := listener.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
