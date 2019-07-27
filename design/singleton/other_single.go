package singleton

import "sync"

//传统方式实现单例模式
var (
	instance *Instance
	lock     sync.Mutex
)

type Instance struct {
	Name string
}

// 双重检查
func GetInstance(name string) *Instance {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Instance{Name: name}
		}
	}
	return instance
}
