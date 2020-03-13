package lru

import (
	"sync"
)

// 存储数据结构体
type Node struct {
	key   interface{}
	value interface{}
	prev  *Node // 往 first 方向
	next  *Node // 往 last 方向
}

// 实现了LRU的结构体
type LRUList struct {
	mux   sync.Mutex
	len   int   // 当前长度
	cap   int   // 最大容量
	first *Node // 队首（最右边），最常使用的
	last  *Node // 队尾（最左边），最少使用的
	nodes map[interface{}]*Node
}

func NewLRUCache(capacity int) *LRUList {
	return &LRUList{
		len:   0,
		cap:   capacity,
		first: nil,
		last:  nil,
		nodes: make(map[interface{}]*Node, capacity),
	}
}

func (l *LRUList) Get(key interface{}) interface{} {
	if node, ok := l.nodes[key]; ok {
		// 将该值设为最后一个
		l.moveToFirst(node)
		return node.value
	}
	return nil
}

func (l *LRUList) Put(key, value interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if node, ok := l.nodes[key]; ok {
		// 更新值
		node.value = value
		l.moveToFirst(node)
	} else {
		// 到达最大容量了，删除最后面的值
		if l.len == l.cap {
			delete(l.nodes, l.last.key)
			l.removeLast(node)
		} else {
			l.len++
		}
	}
	node := &Node{
		key:   key,
		value: value,
	}
	l.nodes[key] = node
	l.insertToFirst(node)
}

func (l *LRUList) moveToFirst(node *Node) {
	// 1. 将该节点从链表里面删掉
	switch node {
	case l.first:
		// 队首，不做改变
		return
	case l.last:
		// 队尾，删掉该节点
		l.removeLast(node)
	default:
		// 队中，删掉该节点
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	// 2. 将该节点插入到队首
	l.insertToFirst(node)
}

func (l *LRUList) removeLast(node *Node) {
	if l.last.prev == nil {
		// 双向链表长度等于1
		l.first = nil
	} else {
		// 双向链表长度大于1
		l.last.prev.next = nil
	}
	l.last = l.last.prev
}

func (l *LRUList) insertToFirst(node *Node) {
	if l.last == nil {
		// 空的链表
		l.last = node
	} else {
		node.next = l.first
		l.first.prev = node
	}
	l.first = node
}

func (l *LRUList) Keys() []interface{} {
	var keys []interface{}
	var last = l.last
	for last != nil {
		keys = append(keys, last.key)
		last = last.prev
	}
	return keys
}
