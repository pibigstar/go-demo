package hystrix

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		MaxConcurrentRequests:  10,   // 最大并发量
		Timeout:                3000, // 单次请求超时时间，如果超过3s将认为服务不可用
		RequestVolumeThreshold: 1,    // 10秒内如果出现1次错误就触发熔断
		ErrorPercentThreshold:  1,    // 按百分比，如果出现1%的错误就触发熔断
		SleepWindow:            2000, // 熔断后2秒后再去尝试服务是否可用
	})

	for i := 0; i < 10; i++ {
		GetInfo(i)
		time.Sleep(1 * time.Second)
	}
}

func GetInfo(i int) {
	hystrix.Go("test", func() error {
		// 业务逻辑部分
		fmt.Println("处理业务逻辑...")
		if i%2 == 0 {
			return errors.New("this is error")
		}
		return nil
	}, func(err error) error {
		// 业务逻辑处理失败, 降级处理
		if err == hystrix.ErrCircuitOpen {
			fmt.Println("服务处于熔断中...")
		} else {
			fmt.Println("服务处理失败, err: ", err.Error())
		}
		return err
	})
}
