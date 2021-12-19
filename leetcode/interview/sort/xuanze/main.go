package main

import "fmt"

func main() {
	arrays := []int{1, 10, 5, 2, 6, 7, 10, 8, 9, 4, 3}
	xuanze(arrays)
	fmt.Println(arrays)
}

/**
对给定的数组进行多次遍历，每次均找出最大的一个值的索引
*/
func xuanze(arrays []int) {
	length := len(arrays)
	for i := 0; i < length; i++ {
		// 最大值的下标
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if arrays[j] > arrays[maxIndex] {
				maxIndex = j
			}
		}
		arrays[length-i-1], arrays[maxIndex] = arrays[maxIndex], arrays[length-i-1]
	}
}
