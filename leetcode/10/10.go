package main

import "fmt"

// 硬币排列
// 第一行排一个硬币，第二个排两个，第三个排三个以此类推
// 求 n 个硬币可以完整排列几行

func main() {
	fmt.Println(find(10))
	fmt.Println(find2(10))
}

// 暴力求解
func find(n int) int {
	for i := 1; i <= n; i++ {
		n = n - i
		if n <= i {
			return i
		}
	}
	return 0
}

// 优化，二分求解
func find2(n int) int {
	low := 0
	high := n
	for low <= high {
		mid := low + (high-low)/2
		// 1，2，3，4...mid
		// 1+2+3+4+5....+mid = ((mid+1)*mid) / 2
		costs := ((mid + 1) * mid) / 2
		if costs == n {
			return mid
		}
		if costs > n {
			high = mid - 1
		} else if costs < n {
			low = mid + 1
		}
	}
	return high
}
