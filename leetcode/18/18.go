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
	fmt.Println(root.Val)
	PreorderByRecursion(root.Left)  // 将左节点入栈
	PreorderByRecursion(root.Right) // 将右节点入栈
}
