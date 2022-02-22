package main

import (
	"testing"
)

// 1 -> 2 -> 3 -> 4 -> 5
func getHead() *Node {
	node5 := NewNode(5, nil)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	return node1
}

func getHead2() *Node {
	node7 := NewNode(7, nil)
	node6 := NewNode(6, node7)
	node0 := NewNode(0, node6)
	return node0
}

// 1 -> 2 -> 2 -> 1 -> nil
func getPalindromes() *Node {
	node4 := NewNode(1, nil)
	node3 := NewNode(2, node4)
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

// 单链表反转-递归
func TestReserve2(t *testing.T) {
	head := getHead()
	newNode := reverse2(head)

	for newNode != nil {
		t.Log(newNode.Val)
		newNode = newNode.Next
	}
}

// 单链表反转前N个节点-递归
func TestReserveN(t *testing.T) {
	head := getHead()
	newNode := reverseN(head, 3)
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

// 链表是否是回文
func TestPalindromes(t *testing.T) {
	t.Log(palindromes(getHead()))
	t.Log(palindromes(getPalindromes()))

	t.Log(isPalindromes(getHead()))
	t.Log(isPalindromes(getPalindromes()))
}

func TestMergeList(t *testing.T) {
	l1 := getHead()
	l2 := getHead2()

	l3 := merge(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}

func TestMergeList2(t *testing.T) {
	l1 := getHead()
	l2 := getHead2()

	l3 := merge2(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}
