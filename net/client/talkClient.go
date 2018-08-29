package client

import (
	"bufio"
	"fmt"
	"go-demo/utils/errutil"
	"net"
	"os"
)

func ClientStart() {

	conn, err := net.Dial("tcp", ":9999")
	errutil.Check(err)

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please write message:")
		//从控制台接收信息
		message, err := reader.ReadString('\n')
		errutil.Check(err)

		fmt.Fprint(conn, message+"\n")

		// 从服务端接收信息
		receiveMessage, err := bufio.NewReader(conn).ReadString('\n')
		errutil.Check(err)

		fmt.Println("receive frmo server:", receiveMessage)
	}
}
