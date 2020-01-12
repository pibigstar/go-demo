package main

import (
	"fmt"
	pb "go-demo/base/rpc/lv2"
)

func main() {
	client, err := pb.DialHelloServiceClient(":8000")
	if err != nil {
		panic(err)
	}
	var resp string
	err = client.Hello("demo2", &resp)
	if err != nil {
		fmt.Printf("err: %s \n", err.Error())
	}
	fmt.Println(resp)
}
