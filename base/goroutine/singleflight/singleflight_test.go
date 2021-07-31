package singleflight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"io/ioutil"
	"testing"
	"time"
)

var (
	data      string
	loadGroup singleflight.Group
)

// 使用SingleFlight合并并发请求
// 例如当缓存被击穿时，可以通过singleFlight让一个请求去请求DB即可
// 其他goroutine共享请求DB的结果
func TestSingleFlight(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			if i == 8 {
				// 告诉 group 忘记这个 key。这样一来，之后这个 key 请求会执行 fn
				// 而不是等待前一个未完成的 fn 函数的结果
				loadGroup.Forget("test")
			}
			// shared 表示是否将结果分享给了其他goroutine
			result, err, shared := loadGroup.Do("test", func() (interface{}, error) {
				// 模拟从缓存取
				if data != "" {
					return data, nil
				}
				// 模拟从数据库中取
				fmt.Println(i)
				bs, err := ioutil.ReadFile("test.txt")
				return string(bs), err
			})
			if err != nil {
				t.Error(err)
			}
			fmt.Println(result, shared)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
