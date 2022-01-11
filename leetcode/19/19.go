package main

import (
	"fmt"
	"go-demo/leetcode/common/tree"
)

// 统计树有多少个节点
// 后序遍历，对于每个节点而言
// 节点总数等于它本身+左子树节点个数+右子树节点个数
func countNode(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := countNode(root.Left)
	rightCount := countNode(root.Right)
	return 1 + leftCount + rightCount
}

// 寻找重复的子树（中等）
// https://leetcode-cn.com/problems/find-duplicate-subtrees/
// 1. 记录以我为根节点的这棵二叉树长什么样子
// 2. 其他节点是否存在跟我一样
var result []int                    // 记录重复的根节点值
var treeMap = make(map[string]bool) // 记录该子树是否已存在
func findDuplicateSubTree(root *tree.TreeNode) {
	if root == nil {
		return
	}
	findDuplicateSubTree(root.Left)
	findDuplicateSubTree(root.Right)
	treeStr := drawTree(root)
	if _, ok := treeMap[treeStr]; ok {
		if !isContains(result, root.Val) {
			result = append(result, root.Val)
		}
	} else {
		treeMap[treeStr] = true
	}
}

func isContains(array []int, v int) bool {
	for _, a := range array {
		if a == v {
			return true
		}
	}
	return false
}

// 描述树
func drawTree(root *tree.TreeNode) string {
	if root == nil {
		return "#"
	}
	left := drawTree(root.Left)
	right := drawTree(root.Right)
	return fmt.Sprintf("%s,%s,%d", left, right, root.Val)
}

// // 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	// 交换左右子树
	left := root.Left
	root.Left = root.Right
	root.Right = left

	// 让左右子节点继续翻转它们的子节点
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
