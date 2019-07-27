package regexp

import (
	"regexp"
	"testing"
)

//正则表达式使用
func TestRegexp(t *testing.T) {
	// 手机号检查
	r := regexp.MustCompile("(13|14|15|17|18|19)[0-9]{9}")
	find := r.Find([]byte("13838254613"))

	t.Log(string(find))
}
