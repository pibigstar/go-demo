package goleak

import (
	"testing"

	"go.uber.org/goleak"
)

func leak() {
	ch := make(chan struct{})
	// goroutine泄漏
	go func() {
		ch <- struct{}{}
	}()
}

// 直接执行并不能发现goroutine泄漏问题
func TestLeak(t *testing.T) {
	leak()
}

// goroutine 泄漏检测
func TestGoLeak(t *testing.T) {
	// 校验goroutine是否有泄漏
	defer goleak.VerifyNone(t)
	leak()
}

// 集成到TestMain中
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
