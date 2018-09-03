package main

import (
	"encoding/base64"
	"fmt"
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


func main() {

	fmt.Println(encrypt("hello world"))

	fmt.Println(UnEncrypt("aGVsbG8gd29ybGQ="))

}