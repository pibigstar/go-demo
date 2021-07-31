package semaphore

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"runtime"
	"testing"
	"time"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)                    // worker数量
	sema       = semaphore.NewWeighted(int64(maxWorkers)) // 信号量
	task       = make([]int, maxWorkers*4)                // 任务数，是worker的四倍
)

// 通过信号量可以有效的控制goroutine的数量
func TestSemaphore(t *testing.T) {
	ctx := context.Background()
	for i := range task {
		// 请求一个worker，如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		// 启动worker goroutine
		go func(i int) {
			// 释放一个worker
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}

	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		t.Logf("获取所有的worker失败: %v", err)
	}

	fmt.Println(task)
}
