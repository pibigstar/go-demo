package main

import "fmt"

func main() {
	arrays := []int{1, 10, 5, 2, 6, 7, 10, 8, 9, 4, 3}
	kuaipai(arrays)
	fmt.Println(arrays)
}

/**
快速排序的原理是，首先找到一个数pivot把数组‘平均’分成两组，
使其中一组的所有数字均大于另一组中的数字，
此时pivot在数组中的位置就是它正确的位置。然后，对这两组数组再次进行这种操作
*/
func kuaipai(arrays []int) {
	recursionSort(arrays, 0, len(arrays)-1)

}

func recursionSort(nums []int, left int, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		recursionSort(nums, left, pivot-1)
		recursionSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	for left < right {
		for left < right && nums[left] <= nums[right] {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		for left < right && nums[left] <= nums[right] {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}

	return left
}
