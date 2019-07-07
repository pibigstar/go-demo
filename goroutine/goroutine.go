package goroutinue

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// 超时检测机制
func GetWithTimeout(d time.Duration) (string, error) {
	ret := make(chan string)
	go func() {
		//do something
		time.Sleep(d)
		ret <- "Hello"
	}()
	select {
	case result := <-ret:
		return result, nil
	case <-time.After(time.Millisecond * 100):
		return "", errors.New("time out")
	}
}

// 当我们启动多个任务时，如果有一个任务完成之后即可返回
// 例如：我们通过关键字去百度和谷歌搜索
// 当其中任意一个先拿到结果之后，我们就直接返回那一个
func GetOne() string {
	ret := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret <- fmt.Sprintf("%d : I get the result", i)
		}(i)
	}
	//当ret中一有结果，就直接返回了
	return <-ret
}

// 当所有任务执行完成之后，才将结果返回
func GetAll() string {
	ret := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret <- fmt.Sprintf("%d : I get the result \n", i)
		}(i)
	}
	// 当所有结果拿到之后，组装结果，返回
	result := ""
	for i := 0; i < 10; i++ {
		result += <-ret
	}
	return result
}

// 使用watiGroup实现获取所有结果之后再返回
func GetAllWithGroup() string {
	ret := make(chan string, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ret <- fmt.Sprintf("%d : I get the result \n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	result := ""
	for i := 0; i < 10; i++ {
		result += <-ret
	}
	return result
}

// 取消任务
func CancelTask() {
	//取消任务标识通道
	cancel := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancel chan struct{}) {
			for {
				if isCancel(cancel) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, ":Canceled")
		}(i, cancel)
	}
	//当关闭通道，那么所有的goroutine都可以监听到
	//如果仅仅是往这个通道发信息，那么只会有一个goroutine被取消
	close(cancel)
	time.Sleep(time.Millisecond * 20)
}

// 判断通道是否取消
func isCancel(cancel chan struct{}) bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

// 使用上下文取消全部任务（一般这些任务分上下级的）
// 一个goroutine中又启动了其他的goroutine
func CancelAllTask() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelFunc(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, ":Canceled")
		}(i, ctx)
	}
	//取消全部任务
	cancelFunc()
	time.Sleep(time.Millisecond * 20)
}

func isCancelFunc(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
