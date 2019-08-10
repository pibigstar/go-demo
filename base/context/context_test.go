package context

import (
	"context"
	"fmt"
	"go-demo/utils/seq"
	"math/rand"
	"testing"
	"time"
)

// 带可取消的Context，调用cancel可取消所有Context执行
// 一般cancel都是放在defer里面，通过return，执行defer里面的取消函数
func TestContextWithCancel(t *testing.T) {
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	getIP := func(ctx context.Context) <-chan string {
		ipChan := make(chan string)
		go func() {
			for {
				i := rand.Intn(1000)
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Microsecond * time.Duration(i)):
					t.Log("模拟耗时操作...")
					ipChan <- seq.UUID()
				}
			}
		}()
		return ipChan
	}

	for ip := range getIP(ctx) {
		t.Logf("Get IP:%s \n", ip)
		return
	}
}

// 带Value的Context
func TestContextWithValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "pibigstar")
	for i := 0; i < 2; i++ {
		go PrintKey(ctx)
	}
}

func PrintKey(ctx context.Context) {
	time.Sleep(time.Microsecond * 2)
	fmt.Println(ctx.Value("name"))
}

// 当到达某一时刻，取消执行
func TestContextWithDeadLine(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "time dead line")
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(50*time.Millisecond))
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		t.Log(ctx.Value("name"))
	case <-ctx.Done():
		t.Log("到达deadLine,结束执行")
	}
}

// 带超时的Context，如果超时之后，将会结束执行
func TestContextWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "time dead line")
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		t.Log(ctx.Value("name"))
	case <-ctx.Done():
		t.Log("已超时,结束执行")
	}
}
