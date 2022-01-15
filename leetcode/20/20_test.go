package main

import (
	"go-demo/leetcode/common/tree"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	kthSmallest(tree.GetBstTree(), 2)
	t.Log(result)
}
