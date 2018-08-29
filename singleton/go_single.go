package singleton

import "sync"

var _goself *GoSingleton

type GoSingleton struct {
	Name string
	sync.Once
}

// 使用go 实现单例模式
func GoInstance(name string) *GoSingleton {
	if _self == nil {
		_goself.Once.Do(func() {
			NewInstance(name)
		})
	}

	return _goself

}

func NewInstance(name string) *GoSingleton {
	_goself = new(GoSingleton)
	_goself.Name = name
	return _goself
}
