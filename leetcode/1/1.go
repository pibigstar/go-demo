package main

// 单链表相关

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

// 链表倒数第k个节点
func findFromEnd(head *Node, k int) *Node {
	p1 := head
	p2 := head

	// 先让p1走k步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	// 然后p1,p2一起走
	// 当p1走到末尾时，p2正好走到n-k处
	// 也就是倒数k的位置
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// 单链表中点
func middleNode(head *Node) *Node {
	p1 := head
	p2 := head

	for p1 != nil && p1.Next != nil {
		// p1 走两步，p2走一步
		// 当p1走到终点时p2正好在中点
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return p2
}
