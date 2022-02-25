package main

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
//影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 动态规划
// 时间复杂度O(n)
// 空间复杂度O(n)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 存放前i个能偷到的最大金额
	// dp[5]=10 表示从前5个最大可偷到5个金额
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// 偷第i间能得到的最大金额为：dp[i-2]+nums[i]
		// 不偷第i间能得到的最大金额为： dp[i-1]
		// 那么dp[i] 就是取偷与不偷的最大值即可
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 动态规划+滚动数组
// 考虑到每间房屋的最高总金额只和该房屋的前两间房屋的最高总金额相关，
// 因此可以使用滚动数组，在每个时刻只需要存储前两间房屋的最高总金额。
// 时间复杂度O(n)
// 空间复杂度O(1)
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
