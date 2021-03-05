package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 1. 建立与服务端的链接
// 2. 进行数据收发
// 3. 关闭链接

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Dial failed , err:", err)
		return
	}
	defer conn.Close() // 关闭连接

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input:")
		msg, _ := reader.ReadString('\n') // 读取用户输入
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
	conn.Close()
}
