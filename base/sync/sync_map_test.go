package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	// 开箱即用
	var users sync.Map

	users.Store("pi", 1)
	users.Store("big", 2)
	users.Store("star", 3)

	if u1, ok := users.Load("pi"); ok {
		fmt.Println(u1)
	}

	// 如果存在则value为存在中的，如果不存在，则为66，并将其放入map
	value, exist := users.LoadOrStore("hello", 66)
	fmt.Println(value, exist)

	users.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	users.Delete("pi")
	fmt.Println("=======delete========")
	users.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
