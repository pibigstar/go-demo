package pb

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/rpc"
)

// 定义 rpc server
const (
	HelloServiceName = "rpc/demo/HelloService"
)

type HelloServiceInterface = interface {
	Hello(req string, resp *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func HandleHTTP() {
	rpc.HandleHTTP()
}

// 定义 rpc client
type HelloServiceClient struct {
	client *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func (h *HelloServiceClient) Hello(req string, resp *string) error {
	return h.client.Call(HelloServiceName+".Hello", req, resp)
}

func (h *HelloServiceClient) AsyncHello(req string, resp *string, done chan *rpc.Call) *rpc.Call {
	return h.client.Go(HelloServiceName+".Hello", req, resp, done)
}

func DialHelloServiceClient(address string) (*HelloServiceClient, error) {
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("base/rpc/lv5/ssl/server.crt")
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		RootCAs: certPool,
	}

	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClient(conn)

	return &HelloServiceClient{client: client}, nil
}
