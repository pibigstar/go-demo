package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// 查看当前程序内存分配情况
func readMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes) \n", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {
	readMemStats()

	arrays := make([]int, 8)
	for i := 0; i < 32*1000*1000; i++ {
		arrays = append(arrays, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}
	readMemStats()

	log.Println(" ===> [start gc]")
	runtime.GC() //强制调用gc回收

	log.Println(" ===> [gc Done].")
	readMemStats()

	time.Sleep(3 * time.Second)
}
