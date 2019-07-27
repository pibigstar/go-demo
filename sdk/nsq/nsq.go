package nsq

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/youzan/go-nsq"
)

func init() {
	flag.StringVar(&env, "env", "dev", "the nsq environment")
}

type (
	Message = nsq.Message
)

type Config struct {
	config   *nsq.Config
	lookAddr string
}

var (
	env       string
	pubMu     = &sync.RWMutex{}
	pubMgrs   = make(map[string]*nsq.TopicProducerMgr)
	consumers = &sync.Map{}
)

var DefaultConfig = func() *Config {
	cfg := nsq.NewConfig()
	return &Config{
		config:   cfg,
		lookAddr: "127.0.0.1:4161",
	}
}()

func getPubMgr(topic string) (*nsq.TopicProducerMgr, error) {
	pubMu.RLock()
	if pubMgr, ok := pubMgrs[topic]; ok {
		pubMu.RUnlock()
		return pubMgr, nil
	}
	pubMu.RUnlock()

	pubMu.Lock()
	defer pubMu.Unlock()

	pubMgr, err := nsq.NewTopicProducerMgr([]string{topic}, DefaultConfig.config)
	if err != nil {
		return nil, err
	}
	err = pubMgr.ConnectToNSQLookupd(DefaultConfig.lookAddr)
	if err != nil {
		return nil, err
	}
	pubMgrs[topic] = pubMgr

	return pubMgr, nil
}

func wrapTopic(topic string) string {
	return fmt.Sprintf("%s-%s", topic, env)
}

func Publish(topic string, data []byte) error {
	topic = wrapTopic(topic)
	pubMgr, err := getPubMgr(topic)
	if err != nil {
		return err
	}
	return pubMgr.Publish(topic, data)
}

func Consume(topic, channel string, handlerFunc nsq.HandlerFunc, concurrency int) error {
	topic = wrapTopic(topic)
	consumer, err := nsq.NewConsumer(topic, channel, DefaultConfig.config)
	if err != nil {
		return err
	}
	consumer.AddConcurrentHandlers(handlerFunc, concurrency)
	// set the consumer to map for close
	key := topic + ":" + channel
	consumers.Store(key, consumer)

	return consumer.ConnectToNSQLookupd(DefaultConfig.lookAddr)
}

func Close() {
	// 记录关闭过得mgr，每个pubMgr仅可被关闭一次
	closedPubMgrs := &sync.Map{}
	pubMu.RLock()
	for _, pubMgr := range pubMgrs {
		if _, ok := closedPubMgrs.Load(pubMgr); ok {
			continue
		}
		closedPubMgrs.Store(pubMgr, struct{}{})
		pubMgr.Stop()
	}
	pubMu.RUnlock()

	// close the consumer
	consumers.Range(func(key, value interface{}) bool {
		if consumer, ok := value.(*nsq.Consumer); ok {
			consumer.Stop()
			select {
			case <-consumer.StopChan:
			case <-time.After(time.Second * 60):
				//等待一分钟让其关闭handler
			}
		}
		consumers.Delete(key)
		return true
	})
}
