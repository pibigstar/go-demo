package sync

import (
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"
)

//存储type，可以实现安全存储不会引发panic
type AtomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue() *AtomicValue {
	return &AtomicValue{}
}

func (av *AtomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value cannot be nil")
	}
	// first set value
	if av.v.Load() == nil {
		av.v.Store(v)
		av.t = reflect.TypeOf(v)
		return nil
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("failed to store value,type:%s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *AtomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *AtomicValue) TypeOfValue() reflect.Type {
	return av.t
}
