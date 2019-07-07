package utils

import (
	"crypto/md5"
	"fmt"
)

// md5 加密
func Md5(str string) string {
	byte := []byte(str)
	sum := md5.Sum(byte)
	md5 := fmt.Sprintf("%x", sum)
	return md5
}
