package server

import (
	"bufio"
	"fmt"
	"go-demo/utils/errutil"
	"net"
	"strings"
)

func ServerStart() {

	fmt.Println("Launching server ....")

	listener, err := net.Listen("tcp", ":9999")
	errutil.Check(err)

	conn, err := listener.Accept()
	errutil.Check(err)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		errutil.Check(err)

		fmt.Println("Received message :", message)

		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage + "\n"))
	}

}
