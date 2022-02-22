package main

import "strconv"

// 给你一个字符串表达式s ，实现一个基本计算器来计算并返回它的值
// 思路：用一个栈保存，所有运算最后都是一个个数做加法
func calculate(s string) int {
	// 栈，将解析过的数字入栈
	var stack []int
	// 上一个遇到的运算符
	prevOp := "+"
	// 当前的num值
	var currNum int

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		}

		if c >= '0' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			// n := int(s[i] - '0')
			currNum = currNum*10 + n
		}

		switch prevOp {
		case "+":
			stack = append(stack, currNum)
		case "-":
			// 将其改为负数入栈
			stack = append(stack, -currNum)
		case "*":
			// 将栈顶值与其做计算
			stack[len(stack)-1] = stack[len(stack)-1] * currNum
		case "/":
			stack[len(stack)-1] = stack[len(stack)-1] / currNum
		}

		prevOp = string(rune(s[i]))
		currNum = 0
	}

	var sum int
	for _, v := range stack {
		sum += v
	}
	return sum
}
