package main

import (
	"log"
	"crypto/md5"
	"fmt"
)

/**
	md5 加密
 */
func main() {

	log.Println(Md5("pibigstar"))

}

func Md5(str string) string {

	byte := []byte(str)

	sum := md5.Sum(byte)

	md5 := fmt.Sprintf("%x", sum)

	return md5
}