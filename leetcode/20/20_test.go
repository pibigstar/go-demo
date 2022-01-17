package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	kthSmallest(tree.GetBstTree(), 2)
	t.Log(result)
}

func TestSearchByBST(t *testing.T) {
	t.Log(searchByBST(tree.GetBstTree(), 3).Val)
}
