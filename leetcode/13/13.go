package main

import "fmt"

// 子数组最大平均数
// 给定一个数组，求长度为k且下标连续的子数组的最大平均数

func main() {
	fmt.Println(maxMean([]int{1, 5, 3, 4, 6, 2, 4}, 3))
	fmt.Println(maxMean2([]int{1, 5, 3, 4, 6, 2, 4}, 3))
}

// 滑动窗口，双指针演变
func maxMean(array []int, k int) int {
	if len(array) < k {
		return 0
	}
	low := 0
	high := k - 1

	var max int
	for high < len(array) {

		var count int
		for i := low; i <= high; i++ {
			count += array[i]
		}
		if count > max {
			max = count
		}
		// 滑动，两个指针一起走
		high++
		low++
	}
	return max / k
}

// 滑动窗口，双指针演变
// 优化，每次滑动其实是整个窗口的值的合减去第一个值，再加上下一次滑动的值
func maxMean2(array []int, k int) int {
	if len(array) < k {
		return 0
	}
	var sum int
	// 第一个滑动窗口的值的合
	for i := 0; i < k; i++ {
		sum += array[i]
	}

	max := sum
	for i := k; i < len(array); i++ {
		// 每次滑动对sum的值进行修改
		sum = sum - array[i-k] + array[i]
		if sum > max {
			max = sum
		}
	}
	return max / k
}
