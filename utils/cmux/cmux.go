package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/soheilhy/cmux"
	grpchello "google.golang.org/grpc/examples/helloworld/helloworld"
)

// 通过一个端口注册多个服务
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:6000")
	if err != nil {
		log.Panic(err)
	}

	m := cmux.New(l)
	// grpc listener
	grpcl := m.Match(cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	// websocket listener
	wsl := m.Match(cmux.HTTP1HeaderField("Upgrade", "websocket"))
	// http listener
	httpl := m.Match(cmux.HTTP1Fast())
	// go rpc listener
	rpcl := m.Match(cmux.Any())

	go serveGRPC(grpcl)
	go serveWS(wsl)
	go serveHTTP(httpl)
	go serveRPC(rpcl)

	if err := m.Serve(); err != nil {
		panic(err)
	}
}

type exampleHTTPHandler struct{}

func (h *exampleHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "example http response")
}

func serveHTTP(l net.Listener) {
	s := &http.Server{
		Handler: &exampleHTTPHandler{},
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func EchoServer(ws *websocket.Conn) {
	if _, err := io.Copy(ws, ws); err != nil {
		panic(err)
	}
}

func serveWS(l net.Listener) {
	s := &http.Server{
		Handler: websocket.Handler(EchoServer),
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

type ExampleRPCRcvr struct{}

func (r *ExampleRPCRcvr) Cube(i int, j *int) error {
	*j = i * i
	return nil
}

func serveRPC(l net.Listener) {
	s := rpc.NewServer()
	if err := s.Register(&ExampleRPCRcvr{}); err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			if err != cmux.ErrListenerClosed {
				panic(err)
			}
			return
		}
		go s.ServeConn(conn)
	}
}

type grpcServer struct{}

func (s *grpcServer) SayHello(ctx context.Context, in *grpchello.HelloRequest) (
	*grpchello.HelloReply, error) {

	return &grpchello.HelloReply{Message: "Hello " + in.Name + " from cmux"}, nil
}

func serveGRPC(l net.Listener) {
	grpcs := grpc.NewServer()
	grpchello.RegisterGreeterServer(grpcs, &grpcServer{})
	if err := grpcs.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}
