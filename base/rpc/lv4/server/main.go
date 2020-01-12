package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	pb "go-demo/base/rpc/lv2"
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
	err := pb.RegisterHelloService(&HelloService{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8000", nil)
}
