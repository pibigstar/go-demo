package main

import "fmt"

// 环形链表
// 判断该链表中是否存在环

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}

func main() {
	node5 := NewNode(5, nil)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)

	node5.Next = node3

	fmt.Println(hasCycle(node1))
	fmt.Println(hasCycle2(node1))
}

// 只要一个元素出现两次，那么必定是有环的
func hasCycle(head *Node) bool {
	exist := make(map[*Node]bool)
	for head != nil {
		if _, ok := exist[head]; ok {
			return true
		}
		exist[head] = true
		head = head.Next
	}
	return false
}

// 使用双指针，优化空间复杂度
func hasCycle2(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	quick := head.Next
	for slow != quick {
		if quick == nil || quick.Next == nil {
			return false
		}
		// 快指针要比慢指针走快点
		quick = quick.Next.Next
		slow = slow.Next
	}
	return true
}
