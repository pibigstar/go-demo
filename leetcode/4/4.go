package main

import "fmt"

// 删除数组中重复元素，返回删除后新数组长度
// 数组是个排序好的数组，要求不能创建新的数组
func main() {
	r := filter([]int{1, 2, 2, 3, 3, 3, 4})
	fmt.Println(r)
}

// 使用双指针算法（快慢指针）
func filter(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int // i是慢指针，j是快指针
	for j := 1; j < len(nums); j++ {
		// 当两者不相同时移动i，j指针
		// 当两者相同时只移动 j 指针
		if nums[i] != nums[j] {
			i++
			// 交换两者的值
			// 也就是将不重复的值往前移
			nums[i] = nums[j]
		}
	}
	newNums := nums[0 : i+1]
	fmt.Println(newNums)
	return i + 1
}
