package pool

import (
	v3 "gopkg.in/go-playground/pool.v3"
)

// 批量执行线程池
type BatchGoPool struct {
	pool    v3.Pool
	batch   v3.Batch
	results chan BatchResult
}

type BatchResult struct {
	Value interface{}
	Error error
}

func NewBatchPool(workers uint) BatchGoPool {
	var pool v3.Pool
	if workers == 0 {
		pool = v3.New()
	} else {
		pool = v3.NewLimited(workers)
	}

	batch := pool.Batch()

	return BatchGoPool{
		pool:    pool,
		batch:   batch,
		results: make(chan BatchResult),
	}
}

func (b *BatchGoPool) Queue(fn func() (interface{}, error)) {
	workFn := func(wu v3.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		return fn()
	}
	b.batch.Queue(workFn)
}

func (b *BatchGoPool) QueueWithArgs(fn func(args ...interface{}) (interface{}, error), args ...interface{}) {
	workFn := func(wu v3.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		return fn(args...)
	}
	b.batch.Queue(workFn)
}

func (b *BatchGoPool) Results() <-chan BatchResult {
	go func(bp *BatchGoPool) {
		for result := range bp.batch.Results() {
			err := result.Error()
			value := result.Value()
			bp.results <- BatchResult{Value: value, Error: err}
		}
		close(bp.results)
	}(b)
	return b.results
}

func (b *BatchGoPool) Cancel() {
	b.batch.Cancel()
}

func (b *BatchGoPool) WaitAll() {
	b.batch.WaitAll()
}

func (b *BatchGoPool) Close() {
	b.pool.Close()
}
