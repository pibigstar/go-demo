package main

import "fmt"

// 最长连续递增序列
// 给定一个未排序的数组，找到最长且连续递增的子序列，并返回序列长度

func main() {
	fmt.Println(maxLength([]int{1, 2, 4, 2, 3, 4, 1, 2, 3, 4, 5}))
}

// 贪心算法
// 找到所有递增子序列的长度，取长度最长的
func maxLength(array []int) int {
	var (
		start int
		max   int
	)
	for i := 1; i < len(array); i++ {
		// 当不是递增时，将start赋值到当前i位置
		if array[i] <= array[i-1] {
			start = i
		}
		// 求出当前子序列的长度
		length := i - start + 1
		// 如果当前子序列长度大于最大值
		// 则将当前子序列的长度设为最大值
		if length > max {
			max = length
		}
	}
	return max
}
