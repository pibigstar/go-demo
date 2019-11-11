package sync

import (
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	locker   sync.Mutex
	rwLocker sync.RWMutex
)

/**
排它锁sync.Mutex
只能加一次锁，重复加锁会导致死锁
*/
func TestMutex(t *testing.T) {
	go mutex("1")
	go mutex("2")
	time.Sleep(time.Millisecond * 10)
}

// 读写锁
func TestRWMutex(t *testing.T) {
	go rwMutex("1")
	go rwMutex("2")
	time.Sleep(time.Millisecond * 10)
}

func mutex(str string) string {
	log.Printf("first:%s", str)
	locker.Lock()
	log.Printf("this is test %s", str)
	if str == "1" {
		return "1"
	}
	time.Sleep(1 * time.Second) //模拟数据库耗时操作
	os.Exit(0)                  //直接退出，验证2 是否执行了
	locker.Unlock()
	return "2"
}

func rwMutex(str string) string {
	log.Printf("first:%s", str)
	rwLocker.RLock() // 加读锁
	rwLocker.RLock() // 可重复加读锁

	//rwLocker.Lock() // 加写锁，只能加一次,一个goroutine中加过之后，另一个再加锁会抛出恐慌

	log.Printf("this is test %s", str)
	if str == "1" {
		return "1"
	}
	time.Sleep(1 * time.Second) //模拟数据库耗时操作
	defer rwLocker.RUnlock()    // 释放读锁
	defer rwLocker.RUnlock()

	return "2"
}
