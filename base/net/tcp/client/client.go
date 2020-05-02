package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please write your msg....")
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from terminal, err: %v \n", err)
			break
		}
		data = strings.TrimSpace(data)

		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("failed to write server: %v \n", err)
		}
		if data == "exit" {
			break
		}

		var result [1024]byte
		n, err := conn.Read(result[:])
		if err != nil {
			fmt.Printf("failed to read result from server, err: %v \n", err)
			continue
		}
		fmt.Printf("read result from server: %s \n", string(result[:n]))
	}
}
