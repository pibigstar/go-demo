package main

import (
	"fmt"
	pb "go-demo/base/rpc/lv5"
)

func main() {
	client, err := pb.DialHelloServiceClient("localhost:1234")
	if err != nil {
		panic(err)
	}
	var result string
	err = client.Hello("test", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// 异步调用
	call := client.AsyncHello("async test", &result, nil)
	<-call.Done
	fmt.Println(result)
}
