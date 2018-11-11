package main

import (
	"log"
	"os"
	"sync"
	"time"
)

/**
读写锁sync.Mutex
 */

 var rwLocker sync.RWMutex
func main() {
	go TestSysncRw("1")
	go TestSysncRw("2")
	select {
	}
}

func TestSysncRw(str string) string {
	log.Printf("first:%s",str)
	rwLocker.RLock() // 加读锁
	rwLocker.RLock() // 可重复加读锁

	//rwLocker.Lock() // 加写锁，只能加一次,一个goroutine中加过之后，另一个再加锁会抛出恐慌

	log.Printf("this is test %s",str)
	if str == "1" {
		return "1"
	}
	time.Sleep(1 * time.Second) //模拟数据库耗时操作
	os.Exit(0) //直接退出，验证2 是否执行了
	rwLocker.RUnlock() // 释放读锁
	rwLocker.RUnlock()

	//rwLocker.Unlock() // 释放写锁
	return "2"
}
