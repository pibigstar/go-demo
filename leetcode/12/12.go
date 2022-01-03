package main

import (
	"fmt"
	"sort"
)

// 合并两个有序数组
// 返回数组1
func main() {
	array1 := []int{1, 2, 3, 7, 9, 0, 0, 0}
	array2 := []int{4, 5, 10}
	fmt.Println(merge(array1, array2))
	fmt.Println(merge2(array1, array2, 5, 3))
}

func merge(array1 []int, array2 []int) []int {
	array1 = append(array1, array2...)
	sort.Ints(array1)
	return array1
}

// 双指针，倒序遍历
// 节省空间复杂度，不需要重新创建数组
func merge2(array1 []int, array2 []int, m, n int) []int {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// 倒序遍历
	for p1 >= 0 && p2 >= 0 {
		// 当array2最后一个值较大时，将其挪到array1 p的位置
		// 并且挪动p2的指针
		if array1[p1] < array2[p2] {
			array1[p] = array2[p2]
			p2--
			p--
		} else {
			// 当array1最后一个值较大时，将其挪到array1 p的位置
			// 并且挪动p1的指针
			array1[p] = array1[p1]
			p1--
			p--
		}
	}
	return array1
}
