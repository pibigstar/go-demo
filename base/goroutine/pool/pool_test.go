package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool, err := NewPool(10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		pool.Put(&Task{
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
			Params: []interface{}{i},
		})
	}

	time.Sleep(time.Second * 1)
}
