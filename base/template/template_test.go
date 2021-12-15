package template

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	m := make(map[string]string)
	m["Hello"] = "你好"
	m["Test"] = "测试"
	m["Error"] = "错误"
	bs, err := gen(m)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bs))
}
