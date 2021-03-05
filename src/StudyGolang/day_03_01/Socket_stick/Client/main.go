package main

// socket_stick/client/main.go

import (
	proto "StudyGolang/day_03_01/Socket_stick/Protocol"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed , err:", err)
			return
		}
		conn.Write(b)
	}
}
