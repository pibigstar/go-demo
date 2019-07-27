package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	Redis.SSet("test", "pibigstar")
	fmt.Println(Redis.SGet("test"))
}
