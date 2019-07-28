package error

import (
	"fmt"
	"runtime"
)

// 打印错误栈
func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}
