package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 监听UDP
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8001,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("upd server listening...")

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		// 回复消息
		go func() {
			str := string(data[:n])
			fmt.Printf("read data: %s, addr: %v, count: %d \n", str, addr, n)

			n, err = listen.WriteToUDP([]byte(fmt.Sprintf("%s OK!", str)), addr)
			if err != nil {
				fmt.Printf("failed to write udp, addr: %v \n", addr)
			}
		}()
	}
}
