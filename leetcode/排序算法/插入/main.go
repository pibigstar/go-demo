package main

import "fmt"

func main() {
	arrays := []int{1, 10, 5, 2, 6, 7, 10, 8, 9, 4, 3}
	charu(arrays)
	fmt.Println(arrays)
}

/**
插入排序的原理是，从第二个数开始向右侧遍历，'
每次均把该位置的元素移动至左侧，放在放在一个正确的位置（比左侧大，比右侧小）。
*/
func charu(arrays []int) {
	for i := 1; i < len(arrays); i++ {
		if arrays[i] < arrays[i-1] {
			j := i - 1
			temp := arrays[i]
			for j >= 0 && arrays[j] > temp {
				arrays[j+1] = arrays[j]
				j--
			}
			arrays[j+1] = temp
		}
	}
}
