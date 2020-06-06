package pb

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义 rpc server
const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(req string, resp *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// 定义 rpc client
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func (h *HelloServiceClient) Hello(req string, resp *string) error {
	return h.Call(HelloServiceName+".Hello", req, resp)
}

func HandleRPCHTTP(w http.ResponseWriter, r *http.Request) {
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}

	err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	if err != nil {
		panic(err)
	}
}
