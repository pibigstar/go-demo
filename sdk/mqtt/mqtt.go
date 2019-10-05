package mqtt

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	Host     = "192.168.1.101:8000"
	UserName = "pibigstar"
	Password = "123456"
)

type Client struct {
	nativeClient  gomqtt.Client
	clientOptions *gomqtt.ClientOptions
	locker        *sync.Mutex
	// 消息收到之后处理函数
	observer func(c *Client, msg *Message)
}

type Message struct {
	ClientID string `json:"clientId"`
	Type     string `json:"type"`
	Data     string `json:"data,omitempty"`
	Time     int64  `json:"time"`
}

func NewClient(clientId string) *Client {
	clientOptions := gomqtt.NewClientOptions().
		AddBroker(Host).
		SetUsername(UserName).
		SetPassword(Password).
		SetClientID(clientId).
		SetCleanSession(false).
		SetAutoReconnect(true).
		SetKeepAlive(120 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetWriteTimeout(10 * time.Second).
		SetOnConnectHandler(func(client gomqtt.Client) {
			// 连接被建立后的回调函数
			fmt.Println("Mqtt is connected!", "clientId", clientId)
		}).
		SetConnectionLostHandler(func(client gomqtt.Client, err error) {
			// 连接被关闭后的回调函数
			fmt.Println("Mqtt is disconnected!", "clientId", clientId, "reason", err.Error())
		})

	nativeClient := gomqtt.NewClient(clientOptions)

	return &Client{
		nativeClient:  nativeClient,
		clientOptions: clientOptions,
		locker:        &sync.Mutex{},
	}
}

func (client *Client) GetClientID() string {
	return client.clientOptions.ClientID
}

func (client *Client) Connect() error {
	return client.ensureConnected()
}

// 确保连接
func (client *Client) ensureConnected() error {
	if !client.nativeClient.IsConnected() {
		client.locker.Lock()
		defer client.locker.Unlock()
		if !client.nativeClient.IsConnected() {
			if token := client.nativeClient.Connect(); token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
	}
	return nil
}

// 发布消息
// retained: 是否保留信息
func (client *Client) Publish(topic string, qos byte, retained bool, data []byte) error {
	if err := client.ensureConnected(); err != nil {
		return err
	}

	token := client.nativeClient.Publish(topic, qos, retained, data)
	if err := token.Error(); err != nil {
		return err
	}

	// return false is the timeout occurred
	if !token.WaitTimeout(time.Second * 10) {
		return errors.New("mqtt publish wait timeout")
	}

	return nil
}

// 消费消息
func (client *Client) Subscribe(observer func(c *Client, msg *Message), qos byte, topics ...string) error {
	if len(topics) == 0 {
		return errors.New("the topic is empty")
	}

	if observer == nil {
		return errors.New("the observer func is nil")
	}

	if client.observer != nil {
		return errors.New("an existing observer subscribed on this client, you must unsubscribe it before you subscribe a new observer")
	}
	client.observer = observer

	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = qos
	}
	client.nativeClient.SubscribeMultiple(filters, client.messageHandler)

	return nil
}

func (client *Client) messageHandler(c gomqtt.Client, msg gomqtt.Message) {
	if client.observer == nil {
		fmt.Println("not subscribe message observer")
		return
	}
	message, err := decodeMessage(msg.Payload())
	if err != nil {
		fmt.Println("failed to decode message")
		return
	}
	client.observer(client, message)
}

func decodeMessage(payload []byte) (*Message, error) {
	message := new(Message)
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.UseNumber()
	if err := decoder.Decode(&message); err != nil {
		return nil, err
	}
	return message, nil
}

func (client *Client) Unsubscribe(topics ...string) {
	client.observer = nil
	client.nativeClient.Unsubscribe(topics...)
}
