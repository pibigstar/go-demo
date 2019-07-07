package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// 存放客户端连接切片
var clients []net.Conn

var (
	host   = "localhost"
	port   = "8888"
	remote = fmt.Sprintf("%s:%s", host, port)
	data   = make([]byte, 1024)
)

func main() {

	fmt.Print("Initiating server........")
	listener, err := net.Listen("tcp", remote)
	defer listener.Close()
	if err != nil {
		log.Printf("Error when listen:%s \n", err.Error())
		os.Exit(-1)
	}
	fmt.Println("Ok!")

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Printf("Error when get connect:%s \n", err.Error())
		}

		clients = append(clients, conn)

		// 为每一个连接分配一个goroutine
		go func(con net.Conn) {

			fmt.Printf("New Connection : %s \n", con.RemoteAddr())
			// 得到客户端的名字
			length, err := con.Read(data)
			if err != nil {
				log.Printf("client : %s quit", con.RemoteAddr())
				con.Close()
				return
			}
			name := string(data[:length])
			connName := name + " enter the room"
			// 通知其他客户端 我进来了
			notify(con, connName)

			// 监听其他客户端的消息
			for {
				length, err := con.Read(data)
				if err != nil {
					log.Printf("client : %s quit", con.RemoteAddr())
					con.Close()
					disconnect(con, name)
					return
				}
				res := string(data[:length])
				msg := fmt.Sprintf("%s said: %s", name, res)
				fmt.Println(msg)

				res = fmt.Sprintf("You said:%s", res)
				con.Write([]byte(res))
				notify(con, msg)
			}
		}(conn)
	}
}

// 群发其他客户端
func notify(conn net.Conn, msg string) {
	for _, con := range clients {
		if con.RemoteAddr() != conn.RemoteAddr() {
			con.Write([]byte(msg))
		}
	}
}

// 断开连接，并通知其他客户端
func disconnect(conn net.Conn, name string) {
	for index, con := range clients {
		if con.RemoteAddr() == conn.RemoteAddr() {
			// 移除此 客户端_
			clients = append(clients[:index], clients[index+1:]...)

			disMsg := fmt.Sprintf("%s %s", name, "left the room")

			notify(con, disMsg)
		}
	}
}
