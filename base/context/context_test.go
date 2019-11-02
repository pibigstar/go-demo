package context

import (
	"testing"
	"time"
)


func TestContextWithCancel(t *testing.T) {
	contextWithCancel()
}

func TestContextWithValue(t *testing.T) {
	kv := make(map[string]interface{})
	kv["name"] = "pibigstar"
	kv["age"] = 20
	contextWithValue(kv)
}

func TestContextWithDeadLine(t *testing.T) {
	contextWithDeadLine(time.Now().Add(50*time.Millisecond))
}

func TestContextWithTimeout(t *testing.T) {
	contextWithTimeout(50*time.Millisecond)
}
