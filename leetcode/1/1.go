package main

import "go-demo/leetcode/common/list"

// 链表反转
func reverse(head *list.Node) *list.Node {
	var prev *list.Node
	var next *list.Node
	for head != nil {
		// 保存当前节点的下一个节点
		next = head.Next
		// 反转当前节点
		head.Next = prev
		// 将当前节点保存为prev
		prev = head
		// 进行下一节点设置
		head = next
	}
	return prev
}

// 链表反转-递归版
func reverse2(head *list.Node) *list.Node {
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
var successor *list.Node

// 反转链表前N个节点-递归版
func reverseN(head *list.Node, n int) *list.Node {
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
func findFromEnd(head *list.Node, k int) *list.Node {
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
func middleNode(head *list.Node) *list.Node {
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

// 判断链表是否是回文链表
func palindromes(head *list.Node) bool {
	// 因为链表没法从后往前遍历，所以我们可以先将链表反转一下
	last := reverse(head)

	for head != nil {
		if head.Val != last.Val {
			return false
		}
		head = head.Next
		last = last.Next
	}
	return true
}

// 判断链表是否是回文链表
// 后续遍历法，树本质就是一个特殊的链表罢了
var left *list.Node

func isPalindromes(head *list.Node) bool {
	temp := *head
	left = &temp
	return palindromes2(head)
}

func palindromes2(right *list.Node) bool {
	if right == nil {
		return true
	}
	// 将节点入栈
	b := palindromes2(right.Next)
	// 出栈时左右进行比较
	b = b && right.Val == left.Val
	left = left.Next
	return b
}

// 合并两个有序链表
func merge(l1 *list.Node, l2 *list.Node) *list.Node {
	head := &list.Node{} // 虚拟头节点
	pre := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}
	return head.Next
}

// 合并两个有序链表， 递归版
func merge2(l1 *list.Node, l2 *list.Node) *list.Node {
	// 当l1为空时递归结束，直接将l2链接到Next上即可
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = merge2(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge2(l1, l2.Next)
		return l2
	}
}

// 环形链表
// 判断该链表中是否存在环
// 只要一个元素出现两次，那么必定是有环的
func hasCycle(head *list.Node) bool {
	exist := make(map[*list.Node]bool)
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
func hasCycle2(head *list.Node) bool {
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
