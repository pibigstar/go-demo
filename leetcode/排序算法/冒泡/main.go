package main

import "fmt"

func main() {
	arrays := []int{1, 10, 5, 2, 6, 7, 10, 8, 9, 4, 3}
	maopao(arrays)
	fmt.Println(arrays)
}

/**
对给定的数组进行多次遍历，每次均比较相邻的两个数，
如果前一个比后一个大，则交换这两个数。经过第一次遍历之后，
最大的数就在最右侧了；第二次遍历之后，第二大的数就在右数第二个位置了
*/
func maopao(arrays []int) {
	for i := 0; i < len(arrays); i++ {
		for j := 1; j < len(arrays)-i; j++ {
			if arrays[j-1] > arrays[j] {
				arrays[j], arrays[j-1] = arrays[j-1], arrays[j]
			}
		}
	}
}
