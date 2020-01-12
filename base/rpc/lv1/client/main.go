package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	var result string
	err = client.Call("HelloService.Hello", "demo1", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
