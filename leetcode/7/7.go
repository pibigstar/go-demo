package main

import (
	"fmt"
	"math"
	"sort"
)

// 求数组中三个数乘积最大值
// 数组中可能有正数，也可能有负数
func main() {
	fmt.Println(findBySort([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(findMax([]int{1, 2, 3, 4, 5, 6}))
}

// 根据排序求
func findBySort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// 当数组中有正数、负数时的最大值
	max1 := nums[0] * nums[1] * nums[n-1]

	// 当数组中全部是正数，或全部是负数时的最大值
	max2 := nums[n-1] * nums[n-2] * nums[n-3]

	if max1 > max2 {
		return max1
	}
	return max2
}

// 根据上面的排序我们可以知道，只需要找出5个数
// 分别是 最小值、第二小值，最大值，第二大值，第三大值 即可求解
func findMax(nums []int) int {
	var (
		min1 = math.MaxInt32 // 最小值
		min2 = math.MaxInt32 // 第二小值
		max1 = math.MinInt32 // 最大值
		max2 = math.MinInt32 // 第二大值
		max3 = math.MinInt32 // 第三大值
	)
	for _, n := range nums {
		// 如果 n小于最小值
		if n < min1 {
			min2 = min1 // 之前的最小值，变成了第二小
			min1 = n
		} else if n < min2 {
			min2 = n
		}

		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}
	}

	// 当有正有负时
	r1 := min1 * min2 * max1
	// 当全部是正，或全部是负时
	r2 := max1 * max2 * max3
	if r1 > r2 {
		return r1
	}
	return r2
}
