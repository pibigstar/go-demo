package main

import "math"

// 最小覆盖子串
// 滑动窗口
func minWindow(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)

	// 统计t中每个字符出现的次数
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 判断need中的字符是否都在当前window中出现过
	// 并且出现次数一致
	// window中出现字符的次数可能会大于need中的，但不能小于，小于肯定不满足
	check := func() bool {
		for k, v := range need {
			// 如果字符k出现的次数少于need中的次数
			// 那么该窗口肯定不满足
			if window[k] < v {
				return false
			}
		}
		return true
	}

	minWinLen := math.MaxInt64
	// 能满足需要的最小左边界，与右边界
	minL, minR := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		// 统计当前窗口中每个字符出现的次数
		window[s[r]]++
		// 如果当前的window窗口满足需求，则开始缩写左边界
		for check() && l <= r {
			// 当前窗口的长度
			currWinLen := r - l + 1

			// 如果当前的窗口长度小于之前的窗口长度
			// 那么就更新最小左右边界
			if minWinLen > currWinLen {
				minL = l
				// 最右边界等于左边界加上窗口的长度
				minR = l + minWinLen
				// 将当前窗口更新为最小窗口长度
				minWinLen = currWinLen
			}

			// window窗口中需要减去一次该字符出现的次数
			window[s[l]]--

			// 右移左边界
			l++
		}
	}

	if minL == -1 {
		return ""
	}
	// 左闭右开
	return s[minL : minR-1]
}
