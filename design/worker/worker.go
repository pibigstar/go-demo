package worker

import "fmt"

// 工人模式
// 自定义工人数量，高效处理任务

var JobQueue chan Job

type Job struct {
	Content string
}

type Worker struct {
	// 工人池
	WorkerPool chan chan Job
	// 任务队列
	JobChan chan Job
	// 退出信号
	QuitChan chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		// WorkerPool: workerPool,
		JobChan:  make(chan Job),
		QuitChan: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChan
			select {
			case job := <-w.JobChan:
				fmt.Println(job.Content)

			case <-w.QuitChan:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// 调度器
type Dispatcher struct {
	WorkPool chan chan Job
	MaxWorks int
}

func NewDispatcher(maxWorks int) *Dispatcher {
	pool := make(chan chan Job, maxWorks)
	return &Dispatcher{WorkPool: pool, MaxWorks: maxWorks}
}

func (d *Dispatcher) Run() {
	// 创建工人
	for i := 0; i < d.MaxWorks; i++ {
		worker := NewWorker(d.WorkPool)
		worker.Start()
	}
	// 调度
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 获取可用的worker通道，如果没有，则一直阻塞
				jobChanel := <-d.WorkPool

				// 将job发送到worker任务通道
				jobChanel <- job
			}(job)
		}
	}
}
