package stack

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

// 获取函数调用栈
func GetStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

// 获取函数调用栈
func GetStackSimple() string {
	stack := debug.Stack()
	return string(stack)
}

// 获取函数调用栈
func GetStackFunc() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("文件: %s, func: %s, 行号: %d", file, funcName, line)
}

// 获取当前go文件的绝对路径
func GetGoFilePath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(currentFile)
	return basePath
}
