package utils

import (
	"encoding/base64"
)

// 使用base64 加密
func encrypt(msg string) string {
	enctry := base64.StdEncoding.EncodeToString([]byte(msg))
	return enctry
}

// base64 解密
func UnEncrypt(encrypt string) string {
	s, _ := base64.StdEncoding.DecodeString(encrypt)
	return string(s)
}

func Base64Encode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	bytes, _ := base64.URLEncoding.DecodeString(str)
	return string(bytes)
}
