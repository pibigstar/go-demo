package deferTest

import (
	"fmt"
	"testing"
	"time"
)

// 巧用 defer, 在函数入口和出口时做一些操作
func TestDefer(t *testing.T) {
	doSomething()
}

func doSomething() {
	defer countTime("doSomething")()
	// 模拟耗时操作
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

// 统计某函数的运行时间
func countTime(msg string) func() {
	start := time.Now()
	fmt.Printf("run func: %s", msg)
	return func() {
		fmt.Printf("func name: %s run time: %f s \n", msg, time.Since(start).Seconds())
	}
}
