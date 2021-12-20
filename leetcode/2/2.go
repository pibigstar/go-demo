package main

import "fmt"

// 字符串切分
// 保证左边与右边字符种类完全相同
// 比如左右有ab字符，那么右边也要有ab字符
// 给定一个不大于10000长度的字符，输出有多少种切分方式

func main() {
	s1 := "abcsabc"    // 无解 0
	s2 := "abcssabc"   // 全对称，1种解
	s3 := "abcssssabc" // 3种
	fmt.Println(findStrCutNum(s1))
	fmt.Println(findStrCutNum(s2))
	fmt.Println(findStrCutNum(s3))
}

// 每个字符 [l1, r1]  l1表示第一次出现位置，r1表示最后一次出现位置
// 为了保证左右字符种类完全相同，那么只能取 [maxL1, minR1] 之间
// 一共有 minR1 - maxL1 种切分方式
func findStrCutNum(s string) int {
	var maxL int      // 保存最大左边界
	var minR = len(s) // 保存最小右边界

	sL := make(map[string]int) // 字符第一次出现的位置
	sR := make(map[string]int) // 字符最后一次出现的位置

	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		str := string(ss[i])
		if _, ok := sL[str]; ok {
			sR[str] = i
		} else {
			sL[str] = i
		}
	}

	flag := 1 // 标识是否有解，当字符出现为1次时那么肯定无解

	// 找出sL中的最大值，sR中的最小值
	for k, v := range sL {
		// 仅出现在左边，无解
		v2, ok := sR[k]
		if !ok {
			flag = 0
			break
		}
		if maxL < v {
			maxL = v
		}
		if minR > v2 {
			minR = v2
		}
	}

	if flag == 0 || minR < maxL {
		return 0
	}
	return minR - maxL
}
