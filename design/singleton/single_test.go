package singleton

import "testing"

// Go语言模式
func TestGoSingleton(t *testing.T) {
	instance1 := GoInstance("pibigstar")
	//查看其内存地址
	t.Logf("%p", instance1)

	instance2 := GoInstance("pibigstar")
	t.Logf("%p", instance2)
}

// 传统模式
func TestSingleton(t *testing.T) {
	instance1 := GetInstance("pibigstar")
	//查看其内存地址
	t.Logf("%p", instance1)

	instance2 := GetInstance("pibigstar")
	t.Logf("%p", instance2)
}
