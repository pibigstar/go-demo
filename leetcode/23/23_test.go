package main

import "testing"

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	t.Log(rob(nums))
	t.Log(rob2(nums))
}
