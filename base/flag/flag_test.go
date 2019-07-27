package flag

import (
	"flag"
	"testing"
)

// 获取命令行参数
// go run flag_demo.go -name pibigstar -age 15
func TestFlag(t *testing.T) {
	// 如果不传，则默认值为 test
	// 如果输入错误参数，那么会打印出 describe
	name := flag.String("name", "pibigstar", "set your name")
	flag.Parse()
	t.Log(*name)
	// 也可以这样使用
	var env string
	flag.StringVar(&env, "env", "dev", "the project environment")
	t.Log(env)
}
