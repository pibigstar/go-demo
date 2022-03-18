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

func TestRegexpDemo(t *testing.T) {
	var text = "Hello 世界！123 Go."
	// 查找连续的小写字母
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["ello" "o"]

	// 查找连续的非小写字母
	reg = regexp.MustCompile(`[^a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["H" " 世界！123 G" "."]

	// 查找连续的单词字母
	reg = regexp.MustCompile(`[\w]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "123" "Go"]

	// 查找连续的非单词字母、非空白字符
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！" "."]

	// 查找连续的大写字母
	reg = regexp.MustCompile(`[[:upper:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["H" "G"]

	// 查找连续的非 ASCII 字符
	reg = regexp.MustCompile(`[[:^ascii:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！"]

	// 查找连续的标点符号
	reg = regexp.MustCompile(`[\pP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["！" "."]

	// 查找连续的非标点符号字符
	reg = regexp.MustCompile(`[\PP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界" "123 Go"]

	// 查找连续的汉字
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界"]

	// 查找连续的非汉字字符
	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello " "！123 Go."]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`Hello|Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找行首以 H 开头，以空格结尾的字符串
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 "]

	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	reg = regexp.MustCompile(`(?i:^hello).*Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 Go"]

	// 查找 Go.
	reg = regexp.MustCompile(`\QGo.\E`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Go."]

	// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^.* `)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以空格开头，到行尾结束，中间不包含空格字符串
	reg = regexp.MustCompile(` [^ ]*$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// [" Go."]

	// 查找“单词边界”之间的字符串
	reg = regexp.MustCompile(`(?U)\b.+\b`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" " 世界！" "123" " " "Go"]

	// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	reg = regexp.MustCompile(`[^ ]{1,4}o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`(?:Hell|G)o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go，替换为 Hellooo、Gooo
	reg = regexp.MustCompile(`(?:Hell|Go)o`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
	// "Hellooo 世界！123 Gooo."

	// 交换 Hello 和 Go
	reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
	// "Go 世界！123 Hello."

	// 特殊字符的查找
	reg = regexp.MustCompile(`[\f\t\n\r\v\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]|]`)
	fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
	// "----------------------"
	// 中文字符
	str := "脚本之家"
	matched, err := regexp.MatchString("[\u4e00-\u9fa5]", str)
	fmt.Println(matched, err)
	// 双字节字符
	str1 := "脚本之家jb51"
	matched1, err := regexp.MatchString("[\u4e00-\u9fa5]", str1)
	fmt.Println(matched1, err)
	// 空字符
	str2 := "\n"
	matched2, err := regexp.MatchString("\\s", str2)
	fmt.Println(matched2, err)
	// email
	str3 := "jb51@163.com"
	matched3, err := regexp.MatchString("\\w[-\\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\\.)+[A-Za-z]{2,14}", str3)
	fmt.Println(matched3, err)
	// 网址
	str4 := "http://www.jb51.net"
	matched4, err := regexp.MatchString("^((https|http|ftp|rtsp|mms)?:\\/\\/)[^\\s]+", str4)
	fmt.Println(matched4, err)
	// 手机号
	str5 := "13688888888"
	matched5, err := regexp.MatchString("0?(13|14|15|18)[0-9]{9}", str5)
	fmt.Println(matched5, err)
	// 国内电话号码
	str6 := "(0516)-88888888"
	matched6, err := regexp.MatchString("[0-9-()（）]{7,18}", str6)
	fmt.Println(matched6, err)
	// 负浮点数
	str7 := "-3.1415926"
	matched7, err := regexp.MatchString("-([1-9]\\d*.\\d*|0.\\d*[1-9]\\d*)", str7)
	fmt.Println(matched7, err)
	// 整数
	str8 := "123456"
	matched8, err := regexp.MatchString("-?[1-9]\\d*", str8)
	fmt.Println(matched8, err)
	// 正浮点数
	str9 := "3.1415926"
	matched9, err := regexp.MatchString("[1-9]\\d*.\\d*|0.\\d*[1-9]\\d*", str9)
	fmt.Println(matched9, err)
	// qq号码
	str10 := "12345678"
	matched10, err := regexp.MatchString("[1-9]([0-9]{5,11})", str10)
	fmt.Println(matched10, err)
	// 邮政编码
	str11 := "221000"
	matched11, err := regexp.MatchString("\\d{6}", str11)
	fmt.Println(matched11, err)
	// ip
	str12 := "192.168.225.255"
	matched12, err := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", str12)
	fmt.Println(matched12, err)
	// 身份证号
	str13 := "320102199002102937"
	matched13, err := regexp.MatchString("\\d{17}[\\d|x]|\\d{15}", str13)
	fmt.Println(matched13, err)
	// 用户名
	str15 := "-jb51"
	matched15, err := regexp.MatchString("[A-Za-z0-9_]+", str15)
	fmt.Println(matched15, err)

}
