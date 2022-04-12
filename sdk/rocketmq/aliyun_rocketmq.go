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

// 使用阿里云的rocketmq

var (
	endpoint   = "http://MQ_INST_1923451467929791_CN8n7179.mq-internet-access.mq-internet.aliyuncs.com:80"
	accessKey  = ""
	secretKey  = ""
	instanceId = ""

	addr, _ = primitive.NewNamesrvAddr(endpoint)
	cre     = primitive.Credentials{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
)

func ConsumerFromAli() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNameServer(addr),
		consumer.WithCredentials(cre),
		consumer.WithNamespace(instanceId),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumerOrder(false), // 是否是顺序消费
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

func PublishFromAli() {
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer(addr),
		producer.WithCredentials(cre),
		producer.WithNamespace(instanceId),
		producer.WithGroupName(group),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ From Ali" + strconv.Itoa(i)),
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
