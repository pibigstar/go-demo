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
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		RequestVolumeThreshold: 1,    // 10秒内如果出现1次错误就触发熔断
		ErrorPercentThreshold:  1,    // 按百分比，如果出现1%的错误就触发熔断
		SleepWindow:            5000, // 熔断后5秒后再去尝试服务是否可用
	})

	GetInfo()

	time.Sleep(3 * time.Second)
}

func GetInfo() {
	hystrix.Go("testHystrix", func() error {
		// 业务逻辑部分
		fmt.Println("处理业务逻辑")
		// 模拟耗时操作，如果超过 配置的 1s,将会进入降级处理
		time.Sleep(1 * time.Second)
		return errors.New("i am dead")
	}, func(err error) error {
		// 业务逻辑处理失败, 降级处理
		fmt.Printf("调用失败, 开始降级处理: %s \n", err.Error())
		return err
	})
}
