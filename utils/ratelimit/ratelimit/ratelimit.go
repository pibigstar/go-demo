package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	// 每秒可以通过100个请求，也就是每个请求间隔10ms
	rl := ratelimit.New(100)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
