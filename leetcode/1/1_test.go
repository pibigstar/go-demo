package main

import (
	"go-demo/leetcode/common/list"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

// 单链表反转
func TestReserve(t *testing.T) {
	head := list.GetHead()
	newNode := reverse(head)

	for newNode != nil {
		t.Log(newNode.Val)
		newNode = newNode.Next
	}
}

// 单链表反转-递归
func TestReserve2(t *testing.T) {
	head := list.GetHead()
	newNode := reverse2(head)

	for newNode != nil {
		t.Log(newNode.Val)
		newNode = newNode.Next
	}
}

// 单链表反转前N个节点-递归
func TestReserveN(t *testing.T) {
	head := list.GetHead()
	newNode := reverseN(head, 3)
	for newNode != nil {
		t.Log(newNode.Val)
		newNode = newNode.Next
	}
}

// 链表倒数第k个节点
func TestFindFromEnd(t *testing.T) {
	head := list.GetHead()
	node := findFromEnd(head, 2)
	t.Log(node.Val)
}

// 单链表中点
func TestMiddleNode(t *testing.T) {
	head := list.GetHead()
	node := middleNode(head)
	t.Log(node.Val)
}

// 链表是否是回文
func TestPalindromes(t *testing.T) {
	t.Log(palindromes(list.GetHead()))
	t.Log(palindromes(list.GetPalindromes()))

	t.Log(isPalindromes(list.GetHead()))
	t.Log(isPalindromes(list.GetPalindromes()))
}

// 合并有序链表
func TestMergeList(t *testing.T) {
	l1 := list.GetHead()
	l2 := list.GetHead2()

	l3 := merge(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}

// 合并有序链表，递归版
func TestMergeList2(t *testing.T) {
	l1 := list.GetHead()
	l2 := list.GetHead2()

	l3 := merge2(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}

// 判断链表是否有环
func TestHasCircle(t *testing.T) {
	t.Log(hasCycle(list.GetCycle()))

	// 双指针法
	t.Log(hasCycle2(list.GetCycle()))
}
