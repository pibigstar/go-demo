package others

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
go2中i:  4
go1中i:  5
go1中i:  5
go1中i:  5
go1中i:  5
go1中i:  5
go2中i:  0
go2中i:  1
go2中i:  2
go2中i:  3
*/

func Test3(t *testing.T) {
	runtime.GOMAXPROCS(1) // 将CPU设置为 1 个
	wg := sync.WaitGroup{}
	wg.Add(10)
	//i是外部for的一个变量，地址不变化。遍历完成后，最终i=5。故go func执行时，i的值始终是5（5次遍历很快完成）。
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("go1中i: ", i) // 5次遍历完之后，i 的值为 5， 而此时goroutine才开始打印
			wg.Done()
		}()
	}

	//i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("go2中i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
