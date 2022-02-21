package main

import (
	"fmt"
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

func TestPrintDFS(t *testing.T) {
	list := printDFS(tree.GetTree())
	for i := 0; i < len(list); i++ {
		vs := list[i]
		fmt.Printf("第%d层 \n", i)
		for _, v := range vs {
			fmt.Printf("%d,", v)
		}
		fmt.Println()
	}
}
