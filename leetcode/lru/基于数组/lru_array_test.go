package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	l := NewLRUArray(10)
	for i := 1; i < 15; i++ {
		l.Put(i)
	}
	fmt.Println(l.arrays)
}
