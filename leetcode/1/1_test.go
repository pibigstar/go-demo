package main

import (
	"testing"
)

func getHead() *Node {
	node5 := NewNode(5, nil)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	return node1
}

func TestMain(m *testing.M) {
	m.Run()
}

// 单链表反转
func TestReserve(t *testing.T) {
	head := getHead()
	newNode := reverse(head)

	for newNode != nil {
		t.Log(newNode.Val)
		newNode = newNode.Next
	}
}

// 链表倒数第k个节点
func TestFindFromEnd(t *testing.T) {
	head := getHead()
	node := findFromEnd(head, 2)
	t.Log(node.Val)
}

// 单链表中点
func TestMiddleNode(t *testing.T) {
	head := getHead()
	node := middleNode(head)
	t.Log(node.Val)
}
