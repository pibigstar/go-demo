package interview

import (
	"fmt"
	"testing"
)

/***
defer函数属延迟执行，延迟到调用者函数执行 return 命令前被执行。多个defer之间按LIFO先进后出顺序执行。
在Panic触发时结束函数运行，在return前先依次打印:打印后、打印中、打印前 。最后由runtime运行时抛出打印panic异常信息。
需要注意的是，函数的return value 不是原子操作.而是在编译器中分解为两部分：返回值赋值 和 return 。
而defer刚好被插入到末尾的return前执行。故可以在derfer函数中修改返回值
*/

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// 发生恐慌，结束执行，结束执行前会先运行defer中的语句
	//panic("发现恐慌")
}

/**
3
2
1
发现恐慌
*/
