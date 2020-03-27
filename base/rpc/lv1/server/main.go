package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// 方法只能有两个可序列化的参数，其中第二个参数是指针类型，
// 并且返回一个error类型，同时必须是公开的方法。
func (*HelloService) Hello(req string, resp *string) error {
	fmt.Printf("Hello: %s \n", req)
	*resp = "Hello: " + req
	return nil
}

func main() {
	// 注册服务
	rpc.RegisterName("HelloService", &HelloService{})

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 调用 rpc server
		go rpc.ServeConn(conn)
	}
}
