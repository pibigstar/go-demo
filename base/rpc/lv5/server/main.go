package main

import (
	"crypto/tls"
	"fmt"
	pb "go-demo/base/rpc/lv5"
	"net/rpc"
)

type HelloService struct {
}

func (*HelloService) Hello(req string, resp *string) error {
	*resp = "this is http rpc: " + req
	return nil
}

func main() {
	err := pb.RegisterHelloService(&HelloService{})
	if err != nil {
		panic(err)
	}

	pb.HandleHTTP()
	fmt.Println("server staring....")

	cert, err := tls.LoadX509KeyPair("base/rpc/lv5/ssl/server.crt", "base/rpc/lv5/ssl/server.key")
	if err != nil {
		panic(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
