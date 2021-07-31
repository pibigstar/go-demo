package encode

import (
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestUTF8ToGBK(t *testing.T) {
	// 使用第三方转码
	enc := mahonia.NewEncoder("gbk")
	s := enc.ConvertString("Hello 派大星")

	// 使用golang官方包转码
	r, err := simplifiedchinese.GBK.NewEncoder().String("Hello 派大星")
	if err != nil {
		t.Error(err)
	}

	if s != r {
		t.Fail()
	}
}
