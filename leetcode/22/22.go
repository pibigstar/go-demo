package main

import "sort"

// 两数之和
func twoSum(arrays []int, target int) [][]int {
	m := make(map[int]int)
	var result [][]int
	for i, a := range arrays {
		s := target - a
		if j, ok := m[s]; ok {
			result = append(result, []int{i, j})
		} else {
			m[a] = i
		}
	}
	return result
}

// 三数之和
// 时间复杂度O(n*n)
// 空间复杂度O(n)
func threeSum(arrays []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(arrays); i++ {
		c := target - arrays[i]
		m := make(map[int]int)
		// 寻找两数之和等于 c 的
		for j := i + 1; j < len(arrays); j++ {
			c1 := c - arrays[j]
			if k, ok := m[c1]; ok {
				result = append(result, []int{arrays[i], arrays[j], arrays[k]})
			}
			m[arrays[j]] = j
		}
	}
	return result
}

// 三数之和，双指值之夹逼法
// 时间复杂度O(n*n)
// 空间复杂度O(1)
func threeSumV2(arrays []int, target int) [][]int {
	var result [][]int
	sort.Ints(arrays)
	for i := 0; i < len(arrays); i++ {
		c := target - arrays[i]

		for j := i + 1; j < len(arrays); j++ {
			k := len(arrays) - 1
			// 夹逼逼近
			for j < k && arrays[j]+arrays[k] > c {
				k--
			}
			// 当两个指针相等时表示没找到
			if j == k {
				break
			}
			if arrays[j]+arrays[k] == c {
				result = append(result, []int{arrays[i], arrays[j], arrays[k]})
			}
		}
	}
	return result
}
