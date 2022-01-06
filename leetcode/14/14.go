package main

import (
	"fmt"
	"go-demo/leetcode/common/tree"
	"math"
)

// 二叉树最小深度

func main() {
	node1 := tree.GetTree()
	// 深度优先
	fmt.Println(minDepth(node1))
	// 广度优先
	fmt.Println(minDepth2(node1))
}

// 深度优先
// 把叶子节点标记为1，从下往上找
func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt32

	// 递归找左节点最小深度
	if root.Left != nil {
		t := minDepth(root.Left)
		if t < min {
			min = t
		}
	}

	// 递归找右节点最小深度
	if root.Right != nil {
		t := minDepth(root.Right)
		if t < min {
			min = t
		}
	}

	return min + 1
}

// 广度优先
// 把root节点标记为1，逐次遍历每层节点
// 维护一个先进先出的队列
func minDepth2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	// 维护一个先进先出的队列
	var queue []*tree.TreeNode
	root.Depth = 1
	queue = append(queue, root)
	for {
		// 出队，先进先出，每次取队首即可
		node := queue[0]
		// 删除刚才取的
		queue = queue[1:]

		// 如果是叶子节点就说明此时已经是最小的深度了
		if node.Right == nil && node.Left == nil {
			return node.Depth
		}

		// 将左右子节点进行入队，并将其深度进行+1
		if node.Left != nil {
			node.Left.Depth = node.Depth + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Depth = node.Depth + 1
			queue = append(queue, node.Right)
		}
	}
}
