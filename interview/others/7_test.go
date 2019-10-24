package others

import (
	"runtime"
	"sync"
	"testing"
)

/**
虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。
map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，
此时同时读写资源存在竞争关系。会报错误信息:“fatal error: concurrent map read and map write”。
*/
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// 压力测试要以Benchmark打头
func BenchmarkMapRW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.GOMAXPROCS(1)
		userAges := &UserAges{}
		userAges.ages = make(map[string]int)
		userAges.Add("A", 1)
		userAges.Get("A")
		userAges.Add("B", 2)
		userAges.GoodGet("B")
	}
}

/**
在执行 Get方法时可能被panic
*/

// 改正后, 加个读锁
// 1. sync.Mutex互斥锁
// 2. sync.RWMutex读写锁，基于互斥锁的实现，可以加多个读锁或者一个写锁。
var rwMutex sync.RWMutex

func (ua *UserAges) GoodGet(name string) int {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
