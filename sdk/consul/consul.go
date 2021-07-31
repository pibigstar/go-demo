package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

var (
	client *api.Client
)

func init() {
	config := &api.Config{
		Address: "127.0.0.1:8500",
	}
	var err error
	client, err = api.NewClient(config)
	if err != nil {
		panic(err)
	}
}

// 注册一个新服务
func Register(id string) {
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    "sns-cnt-center",
		Address: "127.0.0.1",
		Port:    8080,
		Tags:    []string{"sns"},
	}
	//增加check
	check := &api.AgentServiceCheck{
		HTTP:     fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check"),
		Timeout:  "5s", //设置超时 5s
		Interval: "5s", //设置间隔 5s
	}
	//注册check服务
	registration.Check = check

	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println("register server error : ", err)
	}
}

// 读取配置
func ReadConfig(name string) string {
	kvs, _, err := client.KV().List(name, nil)
	if err != nil {
		fmt.Println("get kvs error : ", err)
	}
	var result string
	for _, k := range kvs {
		result = string(k.Value)
	}
	return result
}

// 移除一个服务
func DeleteService(Id string) error {
	return client.Agent().ServiceDeregister(Id)
}

// 服务列表
func ListService() error {
	service, err := client.Agent().Services()
	if err != nil {
		return err
	}
	for k, s := range service {
		fmt.Println(k, s.Address)
	}
	return nil
}
