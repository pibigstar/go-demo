package singleton

import "sync"

//传统方式实现单例模式

var _self *Singleton

type Singleton struct {
	Name string
	sync.Mutex
}

// 双重检查
func Instance(name string) *Singleton {
	if _self == nil {
		_self.Mutex.Lock()         //加锁
		defer _self.Mutex.Unlock() // 在return的时候释放锁
		if _self == nil {
			return &Singleton{Name: name}
		}
	}
	return _self
}
