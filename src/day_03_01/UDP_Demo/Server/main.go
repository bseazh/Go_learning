package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 2000,
	})
	if err != nil {
		fmt.Println("Listen failed , err:", err)
		return
	}
	defer conn.Close()
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("ReadfromUdp failed, err", err)
			return
		}
		fmt.Println("Client Send Message : ", data[:n])
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr) //给Client 回发大写的
	}
}
