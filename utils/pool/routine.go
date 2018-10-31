package pool

import (
	"runtime"

	v3 "gopkg.in/go-playground/pool.v3"
)

var Pool GoPool

type GoPool struct {
	pool v3.Pool
}

func init() {
	poolSize := runtime.NumCPU()
	if poolSize > 1 {
		poolSize = poolSize - 1
	}
	Pool = NewPool(uint(poolSize))
}

func NewPool(workers uint) GoPool {
	var pool v3.Pool
	if workers == 0 {
		pool = v3.New()
	} else {
		pool = v3.NewLimited(workers)
	}
	gp := GoPool{
		pool: pool,
	}
	return gp
}

func (p *GoPool) Queue(fn func()) {
	workFn := func(wu v3.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		fn()

		return nil, nil
	}
	p.pool.Queue(workFn)
}

func (p *GoPool) Go(fn func()) {
	p.Queue(fn)
}

func (p *GoPool) QueueWithArgs(fn func(args ...interface{}), args ...interface{}) {
	workFn := func(wu v3.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		fn(args...)
		return nil, nil
	}
	p.pool.Queue(workFn)
}

func (p *GoPool) GoWithArgs(fn func(args ...interface{}), args ...interface{}) {
	p.QueueWithArgs(fn, args)
}

func (p *GoPool) Reset() {
	p.pool.Reset()
}

func (p *GoPool) Cancel() {
	p.pool.Cancel()
}

func (p *GoPool) Close() {
	p.pool.Close()
}
