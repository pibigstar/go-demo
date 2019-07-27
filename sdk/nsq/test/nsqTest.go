package main

import (
	"fmt"
	"github.com/pibigstar/go-sdk/nsq"
)

func main() {
	err := nsq.Publish("test", []byte("Hello Pibigstar"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c := make(chan struct{})

	nsq.Consume("test", "default", func(msg *nsq.Message) error {
		fmt.Println(string(msg.Body))
		select {
		case <-c:
		default:
			close(c)
		}
		return nil
	}, 5)
	<-c
}
