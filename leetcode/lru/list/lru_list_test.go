package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	l := NewLRUCache(10)
	for i := 1; i < 15; i++ {
		l.Put(i, fmt.Sprintf("第%d个", i))
	}
	for _, key := range l.Keys() {
		fmt.Println(key)
	}
}
