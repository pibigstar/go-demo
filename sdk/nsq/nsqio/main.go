package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var (
	addr          = "127.0.0.1:4150"
	defaultConfig = nsq.NewConfig()
)

// 主函数
func main() {
	Producer("test-dev", []byte("Hello Pibigstar"))

	Consumer("test-dev", "default", HandleMessage)
	time.Sleep(time.Second * 3)
}

func HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

// nsq发布消息
func Producer(topic string, data []byte) error {
	// 新建生产者
	p, err := nsq.NewProducer(addr, defaultConfig)
	if err != nil {
		panic(err)
	}
	// 发布消息
	return p.Publish(topic, data)
}

// 消费消息
func Consumer(topic, channel string, handlerFunc nsq.HandlerFunc) error {
	//新建一个消费者
	c, err := nsq.NewConsumer(topic, channel, defaultConfig)
	if err != nil {
		panic(err)
	}
	//添加消息处理
	c.AddHandler(handlerFunc)
	//建立连接
	return c.ConnectToNSQD(addr)
}

// 运行将会打印： hello NSQ!!!
