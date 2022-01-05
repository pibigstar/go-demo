package main

import (
	"fmt"
	"sort"
)

// 三角形最大周长
// 给定一个数组，找出这个数组中能组成三角形最大周长

func main() {
	fmt.Println(findMaxPerimeter([]int{3, 6, 2, 3}))
}

func findMaxPerimeter(array []int) int {
	// 先排序
	sort.Ints(array)

	// 从后往前找
	for i := len(array) - 1; i >= 2; i-- {
		a := array[i-1]
		b := array[i-2]
		c := array[i]
		// 是否能组成三角形
		if a+b > c {
			return a + b + c
		}
	}
	return 0
}
