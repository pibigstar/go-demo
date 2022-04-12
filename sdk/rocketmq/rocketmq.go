package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strconv"
	"time"
)

// 使用自建的rocketmq

var (
	rocketMQServer = "127.0.0.1:9876"
	topic          = "test"
	group          = "testGroup"
)

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{rocketMQServer})),
	)
	err := c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}
	// Note: start after subscribe
	if err = c.Start(); err != nil {
		panic(err)
	}
	time.Sleep(time.Minute)
	if err = c.Shutdown(); err != nil {
		panic(err)
	}
}

func Publish() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{rocketMQServer})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ" + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	if err = p.Shutdown(); err != nil {
		panic(err)
	}
}
