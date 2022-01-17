package main

import "go-demo/leetcode/common/tree"

// BST 的完整定义如下：
//1、BST 中任意一个节点的左子树所有节点的值都小于该节点的值，右子树所有节点的值都大于该节点的值。
//2、BST 中任意一个节点的左右子树都是 BST

// BST 的中序遍历结果是有序的（升序）

// 寻找BST中第K小的元素
var (
	index  int
	result int
)

func kthSmallest(root *tree.TreeNode, k int) {
	if root == nil {
		return
	}
	kthSmallest(root.Left, k)
	index++
	if index == k {
		result = root.Val
	}
	kthSmallest(root.Right, k)
}

// 在 BST 中搜索元素
func searchByBST(root *tree.TreeNode, target int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	// 去左子树搜索
	if root.Val > target {
		return searchByBST(root.Left, target)
	}
	// 去右子树搜索
	if root.Val < target {
		return searchByBST(root.Right, target)
	}
	return root
}
