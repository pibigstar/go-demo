package pool

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// 简单写一个协程池
var (
	ErrorCapacity   = errors.New("illegal capacity")
	ErrorPoolClosed = errors.New("pool already closed")
)

const (
	START uint8 = 1
	STOP  uint8 = 2
)

type Pool struct {
	// 协程池容量
	capacity int32
	// 工作的协程
	works int32
	// 任务队列
	taskQueue chan *Task
	// 线程池状态
	status uint8
	// 关闭通道
	close chan bool
	// 处理异常
	HandleErr func(interface{})
}

type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

func NewPool(capacity int32) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrorCapacity
	}
	return &Pool{
		capacity:  capacity,
		taskQueue: make(chan *Task, capacity),
		close:     make(chan bool),
		status:    START,
	}, nil
}

// 启动一个work，消费任务队列
func (p *Pool) Run() {
	p.incWorks()

	go func() {
		defer func() {
			p.decWorks()
			if r := recover(); r != nil {
				if p.HandleErr != nil {
					p.HandleErr(r)
				} else {
					fmt.Println("默认异常恢复策略...")
				}
			}
		}()

		// 消费task
		for {
			select {
			case task := <-p.taskQueue:
				task.Handler(task.Params...)
			case <-p.close:
				return
			}
		}
	}()
}

func (p *Pool) Put(task *Task) error {
	if p.status == STOP {
		return ErrorPoolClosed
	}
	if p.works < p.capacity {
		p.Run()
	}
	p.taskQueue <- task
	return nil
}

func (p *Pool) Close() {
	p.status = STOP

	//阻塞等待所有task被消费
	for len(p.taskQueue) > 0 {
	}

	p.close <- true
	close(p.taskQueue)
}

func (p *Pool) incWorks() {
	atomic.AddInt32(&p.works, 1)
}

func (p *Pool) decWorks() {
	atomic.AddInt32(&p.works, -1)
}

func (p *Pool) get() {
	atomic.LoadInt32(&p.works)
}
