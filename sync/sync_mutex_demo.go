package main

import (
  "sync"
  "time"
  "log"
  "os"
)

/**
排它锁sync.Mutex
只能加一次锁，重复加锁会导致死锁
 */
var locker = new(sync.Mutex)
func main() {
  go TestSysnc("1")
  go TestSysnc("2")
  time.Sleep(2 * time.Second)
}

func TestSysnc(str string) string {
  log.Printf("first:%s",str)
  locker.Lock()
  log.Printf("this is test %s",str)
  if str == "1" {
    return "1"
  }
  time.Sleep(1 * time.Second) //模拟数据库耗时操作
  os.Exit(0) //直接退出，验证2 是否执行了
  locker.Unlock()
  return "2"
}