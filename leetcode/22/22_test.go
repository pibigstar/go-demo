package main

import "testing"

func TestTwoSum(t *testing.T) {
	arrays := []int{1, 2, 3, 5, 2, 6}
	result := twoSum(arrays, 8)
	t.Log(result)
}

func TestThreeSum(t *testing.T) {
	arrays := []int{1, 2, 3, 5, 2, 6}
	result := threeSum(arrays, 8)
	t.Log(result)

	result = threeSumV2(arrays, 8)
	t.Log(result)
}
