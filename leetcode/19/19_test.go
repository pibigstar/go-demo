package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

// 统计树有多少个节点
func TestCountNode(t *testing.T) {
	t.Log(countNode(tree.GetTree()))
}

// 寻找重复的子树（中等）
func TestFindDuplicateSubTree(t *testing.T) {
	findDuplicateSubTree(tree.GetHaveDuplicateTree())
	t.Log(result)
}

// 翻转二叉树
func TestInvertTree(t *testing.T) {
	root := tree.GetTree()
	t.Log(drawTree(root))
	// 翻转
	root = invertTree(root)
	t.Log(drawTree(root))
}
