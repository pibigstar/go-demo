package singleton

import "sync"

var (
	goInstance *Instance
	once       sync.Once
)

// 使用go 实现单例模式
func GoInstance(name string) *Instance {
	once.Do(func() {
		goInstance = &Instance{Name: name}
	})
	return goInstance
}
