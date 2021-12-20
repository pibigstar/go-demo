package main

import "fmt"

// 寻找数组中心下标
// 中心下标：该下标左侧的值加一起的合等于右侧值加一起的合
func main() {
	index := findIndex([]int{1, 3, 5, 1, 1, 2})
	fmt.Println(index)
}

// 双指针
// 先计算出总合，慢指针从左向右累加，sum累减
// 当累加值等于累减值，那么该下标即为中心下标
func findIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	var count int
	for i, n := range nums {
		count += n
		if count == sum {
			return i
		}
		sum -= n
	}
	return -1
}
