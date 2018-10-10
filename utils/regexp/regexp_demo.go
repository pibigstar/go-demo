package main

import (
	"fmt"
	"regexp"
)

/**
正则表达式使用
*/
func main() {

	// 手机号检查
	r := regexp.MustCompile("(13|14|15|17|18|19)[0-9]{9}")

	find := r.Find([]byte("13838254613"))
	fmt.Println(string(find))
}
