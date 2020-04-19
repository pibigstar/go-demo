package pb

import (
	"net/rpc"
)

// 定义 rpc server
const HelloServiceName = "rpc/demo/HelloService"

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

func DialHelloServiceClient(address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}
