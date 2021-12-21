package main

import "fmt"

// 两数之和
// 给定一个数组和一个target，返回该数组中
// 两个元素相加等于target的这两个元素的下标
func main() {
	// 暴力循环求解
	fmt.Println(findByBf([]int{2, 7, 11, 15}, 18))
	// 优化，空间换时间
	fmt.Println(findByBf2([]int{2, 7, 11, 15}, 18))
}

// 暴力循环求解
func findByBf(nums []int, target int) (int, int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

// 优化，空间换时间
func findByBf2(nums []int, target int) (int, int) {
	n := len(nums)
	flag := make(map[int]int)
	for i := 0; i < n; i++ {
		if j, ok := flag[target-nums[i]]; ok {
			return j, i
		}
		flag[nums[i]] = i
	}
	return 0, 0
}
