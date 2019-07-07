package goroutinue

import (
	"context"
	"fmt"
	"sync"
	"time"
)
// goroutine超时处理
// 一个请求会触发调用三个服务，每个服务输出一个 int,
// 请求要求结果为三个服务输出 int 之和
// 请求返回时间不超过3秒，大于3秒只输出已经获得的 int 之和
func main() {
	// 超时上下文
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	data := make(chan int,3)
	done := make(chan bool)
	var result int

	// 数据生产者协程
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(3)
		go func() {
			time.Sleep(1 * time.Second)
			data <- 1
			wg.Done()
		}()
		go func() {
			time.Sleep(2 * time.Second)
			data <- 1
			wg.Done()
		}()
		go func() {
			time.Sleep(3 * time.Second)
			data <- 1
			wg.Done()
		}()
		wg.Wait()
	}()

	// 数据消费者协程
	go func() {
		for {
			select {
			// 当data通道有值时
			case d := <-data:
				result += d
				fmt.Println("result",result)
			// 当超过3秒时
			case <-ctx.Done():
				fmt.Println("time out, the result:",result)
				done <- true
			}
		}
	}()

	<-done
}
