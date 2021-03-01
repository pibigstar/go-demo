package sema

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"testing"
	"time"
)

// 通过信号量控制goroutine并发执行的数量
func TestSemaphorePool(t *testing.T) {
	s := semaphore.NewWeighted(1)

	ctx := context.Background()
	for i := 0; i < 20; i++ {
		s.Acquire(ctx, 1)

		go func(i int) {
			defer s.Release(1)
			fmt.Println(i)
			time.Sleep(3 * time.Second)
		}(i)
	}

	err := s.Acquire(ctx, 2)
	if err != nil {
		t.Error(err)
	}
}
