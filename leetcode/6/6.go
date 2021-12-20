package main

import "fmt"

// 求x的平方根，输出最接近的整数

func main() {
	fmt.Println(square(25))
	fmt.Println(square(24))
}

// 二分法暴力求解
func square(x int) int {
	var mid = -1
	var l = 0 // 左边界
	var r = x // 右边界

	for l <= r {
		mid = l + (r-l)/2 // 取 l 与 r 的中间位置
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			l = mid + 1 // 左边界往右挪一格
		} else {
			r = mid - 1 // 右边界往左挪一格
		}
	}
	return mid
}
