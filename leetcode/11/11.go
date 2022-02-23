package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 抢红包算法
// 每次抢到的金额=随机范围[1，M/N *2）
func main() {
	rand.Seed(time.Now().UnixNano())
	// 10元 按分计算
	total := 1000
	// 人数
	num := 10
	for i := 0; i < 10; i++ {
		t := getMoney(total, num)
		fmt.Println(t)
		total -= t
		num--
	}
}

func getMoney(m int, n int) int {
	// 如果就剩一个人，则直接全部返回
	if n == 1 {
		return m
	}
	//区间 [1, 剩余金额/人数的两倍-1)
	t := rand.Intn(m/n*2-1) + 1
	return t
}
