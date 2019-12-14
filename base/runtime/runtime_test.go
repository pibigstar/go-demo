package runtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	//设置线程池中线程最大数等于CPU数，最大的利用CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	go demo1()

	// 打印操作系统
	fmt.Println(runtime.GOOS)
	// 强制GC
	runtime.GC()
}

// 退出协程
func demo1() {
	ticker := time.NewTicker(1)
	<-ticker.C
	// 退出goroutine
	runtime.Goexit()
	fmt.Println("不会执行了")
}
