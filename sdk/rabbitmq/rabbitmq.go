package mq

import (
	"fmt"
	"github.com/pkg/errors"
	"log"

	"github.com/streadway/amqp"
)

var client *amqpClient

type amqpClient struct {
	*amqp.Channel
	done chan bool
}

const (
	// rabbitMq url
	RabbitURL = "amqp://guest:guest@127.0.0.1:5672/"
	//  用于文件transfer的交换机
	TransExchangeName = "pibigstar.trans"
	//  转移队列名
	TransOSSQueueName = "pibigstar.trans.demo"
	//  转移失败后写入另一个队列的队列名
	TransOSSErrQueueName = "pibigstar.trans.demo.err"
	// routingkey
	TransOSSRoutingKey = "demo"
)

func initChannel() bool {
	// 判断channel是否已经被初始化了
	if client != nil {
		return true
	}
	// 获取rabbitmq的一个连接
	conn, err := amqp.Dial(RabbitURL)
	if err != nil {
		fmt.Printf("Failed to dial the rabbitmq,err:%s\n", err.Error())
		return false
	}
	// 打开一个channel，用户消息的发布与接收
	channel, err := conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	client = &amqpClient{Channel: channel}
	return true
}

// put the message to the channel
func (client *amqpClient) PublishText(exchange, routingkey string, msg []byte) error {
	// 判断channel是否正常
	if !initChannel() {
		return errors.New("the channel not initialization")
	}
	pubMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}
	err := client.Publish(exchange, routingkey, false, false, pubMsg)
	if err != nil {
		return err
	}
	return nil
}

// 为了通过CI，这里把init方法给注释掉了
// 如果想运行请记得去掉注释
func init() {
	//initChannel()
}

// monitor the chanel and process the message
func (client *amqpClient) StartConsume(qName, cName string, callback func(msg []byte) bool) {
	// 获取消息信道
	msgs, err := client.Consume(qName, cName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	client.done = make(chan bool)
	// 循环获取队列的消息,为了防止循环一直阻塞代码，用goroutine包裹起来
	go func() {
		for msg := range msgs {
			processSuc := callback(msg.Body)
			if !processSuc {
				//TODO：没有执行成功，写到另一个队列，用于异常情况的重试

			}
		}
	}()
	// done没有新的消息，则会一直阻塞下去
	<-client.done
	log.Printf("close the channel \n")
	// 当有消息过来，也就说明要关闭消费者了
	client.Close()
}

// close the channel
func (client *amqpClient) CloseChannel() {
	client.done <- true
}
