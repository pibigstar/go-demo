package pool

import (
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() interface{} {
		return new(smmall)
	}}

type smmall struct {
	a int
}

//go:noinline
func inc(s *smmall) { s.a++ }

func BenchmarkWithoutPool(b *testing.B) {
	var s *smmall
	for i := 0; i < b.N; i++ {
		s = &smmall{a: 1}
		b.StopTimer()
		inc(s)
		b.StartTimer()
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *smmall
	for i := 0; i < b.N; i++ {
		s = pool.Get().(*smmall)
		s.a = 1
		b.StopTimer()
		inc(s)
		b.StartTimer()
		pool.Put(s)
	}
}
