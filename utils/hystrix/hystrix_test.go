package hystrix

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func TestHystrix(t *testing.T) {
	hystrix.ConfigureCommand("testHystrix", hystrix.CommandConfig{
		// 10s
		Timeout:               10000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	GetInfo()

	time.Sleep(3 * time.Second)
}

func GetInfo() {
	hystrix.Go("testHystrix", func() error {
		// 业务逻辑部分
		fmt.Println("处理业务逻辑")
		// 模拟耗时操作，如果超过 配置的 10s,将会进入降级处理
		time.Sleep(1 * time.Second)
		return errors.New("i am dead")
	}, func(err error) error {
		// 业务逻辑处理失败, 降级处理
		fmt.Printf("调用失败, 开始降级处理: %s \n", err.Error())
		return err
	})
}
