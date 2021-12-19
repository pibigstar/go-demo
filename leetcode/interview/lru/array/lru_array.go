package lru

import "fmt"

type LRUArray struct {
	len    int
	cap    int
	arrays []int
}

func NewLRUArray(capacity int) *LRUArray {
	return &LRUArray{
		cap:    capacity,
		arrays: make([]int, capacity),
	}
}

func (l *LRUArray) Put(value int) {
	index := l.findValue(value)
	if index != -1 {
		l.delete(index)
	} else {
		if l.len >= l.cap {
			l.removeLast()
		}
	}
	l.insertToFirst(value)
}

func (l *LRUArray) findValue(value int) int {
	for i, v := range l.arrays {
		if v == value {
			return i
		}
	}
	return -1
}

func (l *LRUArray) insertToFirst(value int) {
	if l.len > 1 {
		for i := l.len - 1; i > -1; i-- {
			l.arrays[i+1] = l.arrays[i]
		}
	}
	l.arrays[0] = value
	l.len++
}

func (l *LRUArray) removeLast() {
	l.len--
}

func (l *LRUArray) delete(index int) {
	for i := index + 1; i < l.len; i++ {
		l.arrays[i-1] = l.arrays[i]
	}
	fmt.Println("删除元素...")
	l.len--
}
