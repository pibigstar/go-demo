package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	go makeJob()

	dispatcher := NewDispatcher(10)
	dispatcher.Run()

	time.Sleep(time.Second * 1)
}

func makeJob() {
	for i := 0; i < 100; i++ {
		job := Job{Content: fmt.Sprintf("第%d个任务", i)}
		JobQueue <- job
	}
}
