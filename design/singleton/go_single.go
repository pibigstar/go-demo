package singleton

import "sync"

var goSingleton *GoSingleton

type GoSingleton struct {
	Name string
	sync.Once
}

// 使用go 实现单例模式
func GoInstance(name string) *GoSingleton {
	if goSingleton == nil {
		goSingleton.Once.Do(func() {
			NewInstance(name)
		})
	}
	return goSingleton
}

func NewInstance(name string) *GoSingleton {
	goSingleton = new(GoSingleton)
	goSingleton.Name = name
	return goSingleton
}
