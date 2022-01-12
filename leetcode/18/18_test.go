package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

// 前序遍历-递归
func TestPreorderByRecursion(t *testing.T) {
	root := tree.GetTree()
	PreorderByRecursion(root)
}

// 中序遍历-递归
func TestInorderByRecursion(t *testing.T) {
	root := tree.GetTree()
	InorderByRecursion(root)
}

// 后序遍历-递归
func TestAfterOderByRecursion(t *testing.T) {
	root := tree.GetTree()
	AfterOderByRecursion(root)
}
