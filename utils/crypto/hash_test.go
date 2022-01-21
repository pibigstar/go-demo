package utils

import (
	"testing"
)

var data = []byte("Hello")

func TestHash(t *testing.T) {
	t.Log(HashStr(data))
	t.Log(HashNum(data))
}
