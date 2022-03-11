package main

import (
	"fmt"
)

var path []int
var result [][]int

// 找出1...n 之间的所有子集
// 回溯算法
func backtrack(n int, index int) {
	if len(path) > 0 {
		// 每一个路径都是一个子集，加入到结果集中
		result = append(result, append([]int{}, path...))
	}
	// 当起始下标大于n时表示已经到叶子节点了
	if index > n {
		return
	}
	for i := index; i <= n; i++ {
		// 将本次选择的路径加入到选择列表中
		path = append(path, i)
		// 下标+1，进入下一次选择
		backtrack(n, i+1)
		// 回溯
		path = path[:len(path)-1]
	}
}

// 找数组nums的子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		an := append([]int{}, track...)
		fmt.Println(an)
		ans = append(ans, an)
		for i := start; i < len(nums); i++ {
			// 收集子集
			track = append(track, nums[i])
			// 下标+1，同一元素不能重复取
			backtrack(i + 1)
			// 移除最后一个值，回溯
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return ans
}
