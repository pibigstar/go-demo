package sync

import (
	"sync"
	"testing"
)

var (
	mailbox  uint8
	lock     sync.RWMutex
	sendCond = sync.NewCond(&lock)
	recvCond = sync.NewCond(lock.RLocker())
	// sign 用于传递演示完成的信号。
	sign = make(chan struct{}, 2)
	max  = 5
)

// 条件变量，主要起到一个通知的作用
func TestCond(t *testing.T) {
	// 往mailbox放置信息
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			lock.Lock()
			// 如果信箱罗里面有信息了，则让发送方等待
			for mailbox == 1 {
				sendCond.Wait()
			}
			// 信箱没信息，放置信息
			mailbox = 1
			t.Logf("第%d次:放置信息", i)
			lock.Unlock()
			// 通知接收方可以获取信了
			recvCond.Signal()
		}
	}(max)

	// 从mailbox读取信息
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			lock.RLock()
			// 如果信箱没信息，则让接受方等待
			for mailbox == 0 {
				recvCond.Wait()
			}
			mailbox = 0
			t.Logf("第%d次:取走信息", i)
			lock.RUnlock()
			// 通知发送方可以继续放置信了
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
