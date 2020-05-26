package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

var runTimes = 1000

// Use the common pool.
func TestCommonPool(t *testing.T) {
	// 关闭协程池
	defer ants.Release()

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	// 执行100次调用
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		err := ants.Submit(syncCalculateSum)
		if err != nil {
			panic(err)
		}
	}

	wg.Wait()

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}

// Use the pool with a function,
// set 10 to the capacity of goroutine pool and 1 second for expired duration.
func TestPoolWithFunc(t *testing.T) {
	var wg sync.WaitGroup

	// 带方法的协程池
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	// 关闭协程池
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		// 调用协程池中我们定义的方法
		_ = p.Invoke(int32(i))
	}

	wg.Wait()

	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

func TestWithOptions(t *testing.T) {

	p, _ := ants.NewPool(3, ants.WithExpiryDuration(time.Second*1))
	defer p.Release()

	for i := 0; i < 10; i++ {
		err := p.Submit(func() {
			fmt.Printf("执行方法....... %d \n ", i)
			time.Sleep(1 * time.Second) // 模拟耗时操作
			fmt.Println("Hello With Options")
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	time.Sleep(5 * time.Second)
}
