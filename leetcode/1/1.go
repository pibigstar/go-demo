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

// 链表反转-递归版
func reverse2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 后驱节点
// 记录反转之前最后一个节点的后一个节点
// 帮助把反转之后的前N个节点能续上后续节点
var successor *Node

// 反转链表前N个节点-递归版
func reverseN(head *Node, n int) *Node {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
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
