package main

import "fmt"

// 斐波那契数列
// 每一个数都等于前面那两个数之和
// 0，1，1，2，3，5，8，13.....

func main() {
	// 简单递归，暴力求解
	fmt.Println(findByBf(5))
	// 优化，空间换时间
	fmt.Println(findByBf2(5))
}

// 暴力递归求解
func findByBf(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return findByBf(num-1) + findByBf(num-2)
}

// 优化，空间换时间
// 在num-1的递归中已经寻找过了num-2的,没有必要在递归一次
func findByBf2(num int) int {
	arr := make([]int, num+1)
	return findWithArr(arr, num)
}

func findWithArr(arr []int, num int) int {
	if num == 0 || num == 1 {
		return num
	}
	if arr[num] != 0 {
		return arr[num]
	}
	arr[num] = findWithArr(arr, num-1) + findWithArr(arr, num-2)
	return arr[num]
}
