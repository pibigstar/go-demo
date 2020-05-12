package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

//正则表达式使用
func TestRegexp(t *testing.T) {
	// 手机号检查
	r := regexp.MustCompile("(13|14|15|17|18|19)[0-9]{9}")
	find := r.MatchString("13838254613")
	t.Log(find)

	// 正则替换
	// #后添加空格
	str := `#标题1
		    ##标题2`
	r = regexp.MustCompile("#+")
	newStr := r.ReplaceAllStringFunc(str, func(s string) string {
		s = s + " "
		return s
	})
	t.Log(newStr)

	// 分组查询, 判断src里面链接是否为URL
	s := "<img src='https://www.baidu.com' />"
	re := regexp.MustCompile(`<img src=['|"]https.*\.(.*)['|"].*\/>`)
	// 第一个是匹配字符串本身，后面的是正则()里面的内容
	params := re.FindStringSubmatch(s)
	if len(params) > 1 {
		if params[1] != "jpg" {
			fmt.Println(params[1])
		}
	}
}
