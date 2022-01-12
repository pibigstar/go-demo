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
	treeStr := drawTreeAfter(root)
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

// 后续遍历描述树
func drawTreeAfter(root *tree.TreeNode) string {
	if root == nil {
		return "#"
	}
	left := drawTreeAfter(root.Left)
	right := drawTreeAfter(root.Right)
	return fmt.Sprintf("%s,%s,%d", left, right, root.Val)
}

// 前序遍历树描述树
func drawTreePreorder(root *tree.TreeNode) string {
	if root == nil {
		return "#"
	}
	s := fmt.Sprintf("%d", root.Val)
	left := drawTreeAfter(root.Left)
	right := drawTreeAfter(root.Right)
	return fmt.Sprintf("%s,%s,%s", s, left, right)
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

// 二叉树的最大 深度
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	// 求左右子节点的最大深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	var max int
	if leftDepth > rightDepth {
		max = leftDepth
	} else {
		max = rightDepth
	}
	return max + 1
}

// 构造最大二叉树
// https://leetcode-cn.com/problems/maximum-binary-tree/
// 给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
//二叉树的根是数组 nums 中的最大元素。
//左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
//右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
func constructMaximumBinaryTree(nums []int) *tree.TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, low, high int) *tree.TreeNode {
	if low > high {
		return nil
	}
	// 1. 找最大元素作为根
	var (
		max   int
		index int
	)
	for i := low; i <= high; i++ {
		if nums[i] >= max {
			max = nums[i]
			index = i
		}
	}
	// 将max作为根节点
	root := &tree.TreeNode{
		Val: max,
	}
	// 根据index左边值构建左子树
	root.Left = buildTree(nums, low, index-1)
	// 根据index右边值构建右子树
	root.Right = buildTree(nums, index+1, high)
	return root
}

// 根据前序和中序遍历结果，构建二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 前序遍历： 根节点 左子树的值  右子树的值
// 中序遍历： 左子树的值  根节点  右子树的值
// 后序遍历： 左子树的值  右子树的值 根节点
func buildTreeByPreAndInOrder(pre []int, in []int, preStar, preEnd int, inStar, inEnd int) *tree.TreeNode {
	if preStar > preEnd {
		return nil
	}
	// 根据前序结果遍历可知，第一个值为根节点
	rootVal := pre[preStar]

	var index int //根节点在中序遍历中的位置
	for i := inStar; i <= inEnd; i++ {
		if in[i] == rootVal {
			index = i
			break
		}
	}
	root := &tree.TreeNode{
		Val: rootVal,
	}
	leftCount := index - inStar       // 左子树数量
	preLeftEnd := preStar + leftCount // 前序遍历左子树末尾
	// 构建左子树
	root.Left = buildTreeByPreAndInOrder(pre, in, preStar+1, preLeftEnd, inStar, index-1)
	// 构建右子树
	root.Right = buildTreeByPreAndInOrder(pre, in, preLeftEnd+1, preEnd, index+1, inEnd)
	return root
}

// 根据后续遍历和中序遍历构建二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTreeByAfterAndInOrder(after []int, in []int, afterStar, afterEnd int, inStar, inEnd int) *tree.TreeNode {
	if afterStar > afterEnd {
		return nil
	}
	// 根据后序结果遍历可知，最后值为根节点
	rootVal := after[afterEnd]

	var index int //根节点在中序遍历中的位置
	for i := inStar; i <= inEnd; i++ {
		if in[i] == rootVal {
			index = i
			break
		}
	}
	root := &tree.TreeNode{
		Val: rootVal,
	}

	leftCount := index - inStar           // 左子树数量
	afterLeftEnd := afterStar + leftCount // 后续遍历左子树末尾下标
	// 构建左子树
	root.Left = buildTreeByAfterAndInOrder(after, in, afterStar, afterLeftEnd-1, inStar, index-1)
	// 构建右子树
	root.Right = buildTreeByAfterAndInOrder(after, in, afterLeftEnd, afterEnd-1, index+1, inEnd)
	return root
}
