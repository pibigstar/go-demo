package main

import "fmt"

// 寻找100之间有多少个素数（排除0和1）
// 素数：除了1和本身之外没有其他数能整除
func main() {
	// 暴力求解法
	fmt.Println(bf(100))
	// 埃筛法
	fmt.Println(eratosthenes(100))
}

// 暴力算法
func bf(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// 是否是素数
func isPrime(x int) bool {
	// 只需要遍历 根号x即可，也就是 i*i <= x
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 埃筛法
// 素数 * 另一个数 = 非素数，提前把非素数给过滤掉
func eratosthenes(n int) int {
	// true 是非素数
	notPrimes := make([]bool, n)
	var count int
	for i := 2; i < n; i++ {
		// 如果是素数
		if !notPrimes[i] {
			count++
			for j := i * i; j < n; j += i {
				notPrimes[j] = true
			}
		}
	}
	return count
}
