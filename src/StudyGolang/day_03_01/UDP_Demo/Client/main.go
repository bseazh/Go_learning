package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 2000,
	})
	if err != nil {
		fmt.Println("Dialudp failed , err : ", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	var reply [1024]byte
	for {
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("readfromudp failed , err : ", err)
			return
		}
		fmt.Println("收到Server回复消息:", string(reply[:n]))
	}
}
