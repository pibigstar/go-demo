package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

// 统计树有多少个节点
func TestCountNode(t *testing.T) {
	t.Log(countNode(tree.GetTree()))
}

// 二叉树的最大深度
func TestMaxDepth(t *testing.T) {
	t.Log(maxDepth(tree.GetTree()))
}

// 是否是平衡二叉树
func TestIsBalance(t *testing.T) {
	t.Log(isBalance(tree.GetBstTree()))
}

// 寻找重复的子树（中等）
func TestFindDuplicateSubTree(t *testing.T) {
	findDuplicateSubTree(tree.GetHaveDuplicateTree())
	t.Log(result)
}

// 翻转二叉树
func TestInvertTree(t *testing.T) {
	root := tree.GetTree()
	t.Log(drawTreeAfter(root))
	// 翻转
	root = invertTree(root)
	t.Log(drawTreeAfter(root))
}

// 构造最大二叉树
func TestConstructMaximumBinaryTree(t *testing.T) {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 2, 5})
	t.Log(drawTreePreorder(root))
}

// 根据前序和中序遍历结果，构建二叉树
func TestBuildTreeByPreAndInOrder(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTreeByPreAndInOrder(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
	t.Log(drawTreePreorder(root))
}

// 根据后序和中序遍历结果，构建二叉树
func TestBuildTreeByAfterAndInOrder(t *testing.T) {
	afterOrder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTreeByAfterAndInOrder(afterOrder, inorder, 0, len(afterOrder)-1, 0, len(inorder)-1)
	t.Log(drawTreePreorder(root))
}
