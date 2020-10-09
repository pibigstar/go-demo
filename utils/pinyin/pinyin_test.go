package pinyin

import (
	"fmt"
	"testing"

	"github.com/mozillazg/go-pinyin"
)

// 汉字转拼音
func TestPinyin(t *testing.T) {
	str := "派大星"
	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(str, a))
	// [[pai] [da] [xing]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(str, a))
	// [[pài] [dà] [xīng]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(str, a))
	// [[pa4i] [da4] [xi1ng]]
}
