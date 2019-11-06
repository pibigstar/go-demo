package stack

import (
	"runtime"
)

// 获取函数调用栈
func GetStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
