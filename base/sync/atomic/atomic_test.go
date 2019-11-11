package sync

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 原子存储
func TestAtomicValue(t *testing.T) {
	v := "pibigstar"
	av := NewAtomicValue()
	av.Store(v)
	t.Log(av.Load())
}

// 不使用原子操作做加减
func TestNotUserAtomic(t *testing.T) {
	// not use atomic
	var sum = 10
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			sum += t
		}(i)
	}
	wg.Wait()
	t.Log(sum)
}

// 使用原子操作进行加减
func TestUseAtomic(t *testing.T) {
	var sum int32 = 10
	var wg sync.WaitGroup
	for i := int32(0); i < 10; i++ {
		wg.Add(1)
		go func(t int32) {
			defer wg.Done()
			// 确保了同一时间只有一个Goroutine执行该操作
			atomic.AddInt32(&sum, t)
		}(i)
	}
	wg.Wait()
	t.Log(sum)
}

// 采用CSA的原子操作
func TestUseAtomicCAS(t *testing.T) {
	var sum int32 = 10
	var wg sync.WaitGroup
	for i := int32(0); i < 10; i++ {
		wg.Add(1)
		go func(t int32) {
			defer wg.Done()
			// 先比较变量的值是否等于给定旧值，等于旧值的情况下才赋予新值，最后返回新值是否设置成功
			atomic.CompareAndSwapInt32(&sum, sum, sum+t)
		}(i)
	}
	wg.Wait()
	t.Log(sum)
}

func TestSpinLock(t *testing.T) {
	t.Log(SpinLock())
}
