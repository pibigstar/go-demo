package singleton

import "sync"

type singleton struct{}

var (
	goInstance *Instance
	ins        *singleton
	once       sync.Once
)

// 使用go 实现单例模式
func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
