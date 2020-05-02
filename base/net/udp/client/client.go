package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8001,
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		_, err := conn.Write([]byte("Hello World!"))
		if err != nil {
			fmt.Printf("failed to write msg, err: %v \n", err)
			break
		}
		var result [1024]byte
		n, addr, err := conn.ReadFromUDP(result[:])
		if err != nil {
			fmt.Printf("failed to receiver msg, addr: %v, err: %v \n", addr, err)
		}
		fmt.Printf("receiver msg from: %v, msg: %s \n", addr, string(result[:n]))
	}
}
