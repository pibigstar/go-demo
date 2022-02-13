package runtime

import (
	"fmt"
	"os/exec"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
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

func TestSetMaxThread(t *testing.T) {
	// 设置最大线程数
	// 超过这个线程数之后程序会支付fatal
	debug.SetMaxThreads(10)
	for i := 0; i < 2; i++ {
		go func() {
			// 通过执行 shell 命令 sleep 3 来模拟同步系统调用过程
			_, err := exec.Command("bash", "-c", "sleep 1").Output()
			if err != nil {
				panic(err)
			}
		}()
	}
	time.Sleep(time.Second)
}

// 通过runtime.LockOSThread()将M与G进行绑定
// 在不执行runtime.UnlockOSThread的情况下
// 当G退出后，空闲的M也会退出，从而达到释放空闲线程的作用
func TestLockOSThread(t *testing.T) {
	threadProfile := pprof.Lookup("threadcreate")
	t.Logf("初始化时线程数： %d\n", threadProfile.Count())

	for i := 0; i < 3; i++ {
		go func() {
			runtime.LockOSThread()
			_, err := exec.Command("bash", "-c", "sleep 1").Output()
			if err != nil {
				panic(err)
			}
		}()
	}
	time.Sleep(time.Second)
	t.Logf("结束时线程数: %d\n", threadProfile.Count())
}
