package main

import (
	"flag"
	"log"
)

/**
获取命令行参数
go run flag_demo.go -name pibigstar -age 15
*/
func main() {
	// 如果不传，则默认值为 test
	// 如果输入错误参数，那么会打印出 describe
	name := flag.String("name", "test", "describe：set your name")

	age := flag.Int("age", 18, "describe: set your age")

	flag.Parse()
	log.Println(*name)
	log.Println(*age)

}
