package main

import "fmt"

// 链表反转

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

	newNode := reverse(node1)

	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}

// 链表反转
func reverse(head *Node) *Node {
	curr := head
	var prev *Node
	var next *Node

	for curr != nil {
		// 保存当前节点的下一个节点
		next = curr.Next
		// 反转当前节点
		curr.Next = prev
		// 将当前节点保存为prev
		prev = curr
		// 进行下一节点设置
		curr = next
	}
	return prev
}
