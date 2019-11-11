package sync

import (
	"sync/atomic"
	"time"
)

// 利用CompareAndSwapInt32实现一个自旋锁
func SpinLock() int32 {
	var num int32 = 1
	for {
		if atomic.CompareAndSwapInt32(&num, 1, 0) {
			break
		}
		time.Sleep(time.Microsecond * 100)
	}
	return num
}
