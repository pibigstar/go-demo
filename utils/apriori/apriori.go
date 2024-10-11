package main

import (
	"fmt"
)

/**
*  @Author: leikewei
*  @Date: 2024/8/18
*  @Desc: 先验算法
 */

func main() {

	transactions := [][]string{
		{"豆奶", "莴苣"},
		{"豆奶", "莴苣"},
		{"豆奶", "莴苣"},
		{"豆奶", "莴苣"},
		{"豆奶", "莴苣"},
		{"豆奶", "莴苣"},
		{"莴苣", "尿布", "葡萄酒", "甜菜"},
		{"豆奶", "尿布", "葡萄酒", "橙汁"},
		{"莴苣", "豆奶", "尿布", "葡萄酒"},
		{"莴苣", "豆奶", "尿布", "橙汁"},
	}
	// orderedStatistic:
	apriori := Apriori.NewApriori(transactions)
	// minSupport: 最小支持度
	// minConfidence: 最小置信度
	results := apriori.Calculate(Apriori.NewOptions(0.5, 1, 0.0, 0))

	for _, r := range results {
		fmt.Printf("%+v, \n", r)
	}
}
