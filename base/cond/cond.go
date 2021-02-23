package main

import (
	"fmt"
	"sync"
)

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// 消费者
	go func() {
		for i := 0; i < 100; i++ {
			// 消费者开始消费时，锁住
			cond.L.Lock()
			// 如果没有可消费的值，则等待
			for condition == 0 {
				cond.Wait()
			}
			// 消费
			condition--
			fmt.Printf("Consumer: %d\n", condition)

			// 唤醒一个生产者
			cond.Signal()
			// 解锁
			cond.L.Unlock()
		}
	}()

	// 生产者
	for i := 0; i < 100; i++ {
		// 生产者开始生产
		cond.L.Lock()

		// 当生产太多时，等待消费者消费
		for condition == 10 {
			cond.Wait()
		}
		// 生产
		condition++
		fmt.Printf("Producer: %d\n", condition)

		// 通知消费者可以开始消费了
		cond.Signal()
		// 解锁
		cond.L.Unlock()
	}
}
