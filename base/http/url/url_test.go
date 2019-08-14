package url

import (
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	v := url.Values{}
	v.Add("orgId", "123456")
	v.Add("userId", "pibigstar")

	// url编码
	body := v.Encode()
	t.Log(body)
	// url解码
	values, err := url.ParseQuery(body)
	if err != nil {
		t.Error(err)
	}
	t.Log(values.Get("orgId"))
}

func TestPathEscape(t *testing.T) {
	//  将 // ? 这些特殊字符编码
	encode := url.PathEscape("http://www.baidu.com?username=123&password=456")
	t.Log(encode)

	// 将 %2F 解码 为 /
	decode, err := url.PathUnescape(encode)
	if err != nil {
		t.Error(err)
	}
	t.Log(decode)
}
