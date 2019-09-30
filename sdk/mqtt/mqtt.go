package mqtt

import (
	"encoding/json"
	"fmt"
	gomqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"sync"
	"time"
)

var Mqtt *mqtt

type mqtt struct {
	host     string
	username string
	password string
}

func (mqtt *mqtt) Init(conf map[string]interface{}) (err error) {
	mqtt.host = cast.ToString(conf["host"])
	mqtt.username = cast.ToString(conf["username"])
	mqtt.password = cast.ToString(conf["password"])
	return err
}
func (*mqtt) Close() {
}

type MqttClient struct {
	nativeClient  gomqtt.Client
	clientOptions *gomqtt.ClientOptions
	locker        *sync.Mutex
	// 消息收到之后处理函数
	observer func(c *MqttClient, msg *MqttMessage)
}

type MqttMessage struct {
	ClientID   string `json:"clientId"`
	Time       int64  `json:"time"`
	EventType  string `json:"eventType"`
	EventData  int64  `json:"eventData,omitempty"`
	EventIndex int64  `json:"eventIndex,omitempty"`
	ChannelID  string `json:"channelId,omitempty"`
	ReqId      string `json:"reqId,omitempty"`
}

func (mqtt *mqtt) NewClient(clientId string) *MqttClient {
	clientOptions := gomqtt.NewClientOptions().
		AddBroker(mqtt.host).
		SetUsername(mqtt.username).
		SetPassword(mqtt.password).
		SetClientID(clientId).
		SetCleanSession(false).
		SetAutoReconnect(true).
		SetKeepAlive(120 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetWriteTimeout(10 * time.Second).
		SetOnConnectHandler(func(client gomqtt.Client) {
			fmt.Println("Mqtt connected...", "clientId", clientId)
		}).
		SetConnectionLostHandler(func(client gomqtt.Client, err error) {
			fmt.Println("Mqtt disconnected.", "clientId", clientId, "reason", err.Error())
		})

	nativeClient := gomqtt.NewClient(clientOptions)

	return &MqttClient{
		nativeClient:  nativeClient,
		clientOptions: clientOptions,
		locker:        &sync.Mutex{},
	}
}

// mqtt 常用操作
func (client *MqttClient) GetClientID() string {
	return client.clientOptions.ClientID
}

func (client *MqttClient) Connect() error {
	return client.ensureConnected()
}

// 确保连接
func (client *MqttClient) ensureConnected() error {
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

func (client *MqttClient) Publish(topic string, qos byte, retained bool, data []byte) error {
	if err := client.ensureConnected(); err != nil {
		return err
	}

	token := client.nativeClient.Publish(topic, qos, retained, data)
	if err := token.Error(); err != nil {
		return err
	}

	// return false is the timeout occurred
	if !token.WaitTimeout(time.Second * 10) {
		fmt.Println("mqtt publish wait timeout")
	}

	return nil
}

func (client *MqttClient) Subscribe(observer func(c *MqttClient, msg *MqttMessage), qos byte, topics ...string) error {
	if observer == nil {
		return nil
	}
	if client.observer != nil {
		return status.Error(codes.Unavailable, "an existing observer subscribed on this client, you must unsubscribe it before you subscribe a new observer.")
	}
	client.observer = observer
	client.subscribeMultiple(qos, topics...)
	return nil

}

func (client *MqttClient) subscribeMultiple(qos byte, topics ...string) {
	if len(topics) == 0 || client.observer == nil {
		return
	}
	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = qos
	}
	client.nativeClient.SubscribeMultiple(filters, client.mqttMessageHandler)
}

func (client *MqttClient) mqttMessageHandler(c gomqtt.Client, msg gomqtt.Message) {
	if client.observer == nil {
		fmt.Println("not subscribe mqtt message observer")
		return
	}

	mqttMessage, err := decodeMqttMessage(msg.Payload())
	if err != nil {
		fmt.Println("failed to decode incoming message")
		return
	}

	if mqttMessage.ReqId == "" {
		mqttMessage.ReqId = uuid.NewV1().String()
	}
	client.observer(client, mqttMessage)
}

func decodeMqttMessage(payload []byte) (*MqttMessage, error) {
	mqttMessage := new(MqttMessage)
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.UseNumber()
	if err := decoder.Decode(&mqttMessage); err != nil {
		return nil, err
	}
	return mqttMessage, nil
}

func (client *MqttClient) Unsubscribe(topics ...string) {
	client.observer = nil
	client.nativeClient.Unsubscribe(topics...)
}
