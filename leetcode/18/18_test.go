package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

// 二叉树遍历
// 前中后 是根据根节点来的，根节点在前面就是前序遍历
// 前序遍历： 根左右 （先打印根节点，再打印左右子节点）
// 中序遍历： 左根右
// 后序遍历： 左右根
// 层序遍历： 从上往下，从左往右（按层级来的）

// 前序遍历-递归
func TestPreorderByRecursion(t *testing.T) {
	root := tree.GetTree()
	PreorderByRecursion(root)
}
