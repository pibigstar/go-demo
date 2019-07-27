package string

import (
	"log"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "Hello 派大星"
	// 这个是获取 H 的Unicode的
	s := str[0]
	log.Println("str[0]:", s)

	//遍历每一个字符
	for i, j := range str {
		// 注意： i 不是递增的，i 是根据 他的字节进行增加的
		log.Printf("第%d个字符：%s", i, string(j))
	}

	// 这个是获取字节的长度 15 = 6 + 3*3  汉字占3个字节
	l := len(str)
	log.Println("len:", l)

	// 使用（多个）空格分割字符串
	fields := strings.Fields(str)
	log.Println("fields:", fields)

	// 以特殊字符进行分割
	split := strings.Split(str, "o")
	t.Log("split:", split)

	// 查看字符串是否包含 一个字符串
	t.Log("是否包含字符串 llo  :", strings.Contains(str, "llo"))

	// 查看字符串是否以某个字符串开头
	t.Log("是否以字符串 He 开头：", strings.HasPrefix(str, "He"))

	// 查看字符串是否以某个字符串结尾
	t.Log("是否以字符串 星 开头：", strings.HasSuffix(str, "星"))

	// 查看字符串中另一个字符串出现的次数
	t.Log("l 在str字符串中出现了：", strings.Count(str, "l"))

	// 返回字符串中另一个字符串出现的位置
	t.Log("派 在str字符串中第一次出现的位置为：", strings.Index(str, "派"))
	// 字符串替换  n 为替换几次 如果为-1 则全部替换
	t.Log("字符串替换：", strings.Replace(str, "l", "g", -1))

	// 去除字符串首尾空格
	t.Log("去空格：", strings.TrimSpace(str))
	// 去除字符串左边特定字符
	t.Log("去左边：", strings.TrimLeft(str, "He"))
	// 去除字符串右边边特定字符
	t.Log("去左边：", strings.TrimRight(str, "大星"))
	// 去除字符串首位的字符
	t.Log("去首位：", strings.Trim(str, "H"))
}
