package main

import "fmt"

// 柠檬水找零
// 柠檬水为5元，面值只有 5， 10， 20
// 判断最终能否成功找零
func main() {
	fmt.Println(change([]int{5, 10, 20}))
}

func change(bills []int) bool {
	var (
		five int // 5元的个数
		ten  int // 10元的个数
	)
	for _, b := range bills {
		if b == 5 {
			five++
		}
		if b == 10 {
			// 如果此时手里没有5元的，那么肯定没法找零了
			if five == 0 {
				return false
			}
			// 可以成功找零，减去一个5元
			five--
			ten++
		}
		if b == 20 {
			// 局部最优，优先找10+5的
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 { // 只能找3个5元了
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
