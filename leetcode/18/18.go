package main

import (
	"fmt"
	"go-demo/leetcode/common/tree"
)

// 前序遍历-递归
func PreorderByRecursion(root *tree.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)           // 前序：第一次在栈顶时就打印
	PreorderByRecursion(root.Left)  // 将左节点入栈
	PreorderByRecursion(root.Right) // 将右节点入栈
}

// 中序遍历-递归
func InorderByRecursion(root *tree.TreeNode) {
	if root == nil {
		return
	}
	InorderByRecursion(root.Left)  // 将左节点入栈
	fmt.Println(root.Val)          // 中序：第二次在栈顶时就打印
	InorderByRecursion(root.Right) // 将右节点入栈
}

// 后序遍历-递归
func AfterOderByRecursion(root *tree.TreeNode) {
	if root == nil {
		return
	}
	AfterOderByRecursion(root.Left)  // 将左节点入栈
	AfterOderByRecursion(root.Right) // 将右节点入栈
	fmt.Println(root.Val)            // 后序：第三次在栈顶时就打印
}
