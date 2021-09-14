package chanx_test

import (
	"context"
	"testing"
	"time"
)

// 使用chan实现 try lock 功能
type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) UnLock() {
	select {
	case m.ch <- struct{}{}:
	default:
	}
}

// 判断锁是否已经别人持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

// 带超时的 try lock
func (m *Mutex) TryLock(ctx context.Context) bool {
	select {
	case <-m.ch:
		return true
	case <-ctx.Done():
		return false
	}
}

func TestTryLock(t *testing.T) {
	mu := NewMutex()
	t.Log("Lock: ", mu.IsLocked())

	mu.Lock()

	// 10毫秒后释放锁
	go func() {
		time.AfterFunc(10*time.Millisecond, func() {
			mu.UnLock()
		})
	}()

	go func() {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Millisecond)
		if mu.TryLock(ctx) {
			t.Log("加锁成功1")
			mu.UnLock()
		} else {
			t.Log("加锁失败1")
		}
	}()

	go func() {
		// 会超时
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Millisecond)
		if mu.TryLock(ctx) {
			t.Log("加锁成功2")
			mu.UnLock()
		} else {
			t.Log("加锁失败2")
		}
	}()

	time.Sleep(1 * time.Second)
}
