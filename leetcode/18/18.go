package main

import (
	"fmt"
	"go-demo/leetcode/common/tree"
)

// 二叉树遍历
// 前中后 是根据根节点来的，根节点在前面就是前序遍历
// 前序遍历： 根左右 （先打印根节点，再打印左右子节点）
// 中序遍历： 左根右
// 后序遍历： 左右根
// 层序遍历： 从上往下，从左往右（按层级来的）

// 前中后序是遍历二叉树过程中处理每一个节点的三个特殊时间点，绝不仅仅是三个顺序不同的 List
// 前序位置的代码在刚刚进入一个二叉树节点的时候执行；
// 中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。
// 后序位置的代码在将要离开一个二叉树节点的时候执行；

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

var dfsList map[int][]int

func printDFS(root *tree.TreeNode) map[int][]int {
	dfsList = make(map[int][]int)
	dfs(root, 0)
	return dfsList
}

// 层序遍历
func dfs(root *tree.TreeNode, level int) {
	if root == nil {
		return
	}
	dfsList[level] = append(dfsList[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
