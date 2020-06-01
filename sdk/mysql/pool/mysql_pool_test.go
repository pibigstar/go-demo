package pool

import (
	"context"
	"fmt"
	"testing"
)

func TestPool1(t *testing.T) {
	ctx := context.Background()
	config := &Config{
		MaxConn:        2,
		MaxIdle:        1,
		MaxWait:        10,
		MaxWaitTimeout: 3000,
	}
	conn := Prepare(ctx, config)
	if _, err := conn.New(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}
	// 释放连接
	go conn.Release(ctx)

	if _, err := conn.New(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}
	// 释放连接
	go conn.Release(ctx)

	if _, err := conn.New(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := conn.New(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}
}
