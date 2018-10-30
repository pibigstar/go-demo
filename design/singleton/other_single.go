package singleton

import "sync"

//传统方式实现单例模式
var singleton *Singleton

type Singleton struct {
	Name string
	sync.Mutex
}

// 双重检查
func Instance(name string) *Singleton {
	if singleton == nil {
		singleton.Mutex.Lock()         //加锁
		defer singleton.Mutex.Unlock() // 在return的时候释放锁
		if singleton == nil {
			return &Singleton{Name: name}
		}
	}
	return singleton
}
