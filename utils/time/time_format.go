package main

import (
	"time"
	"log"
)

const defaultTime  = "2006-01-02 15:04:05"

// 时间格式化
func TimeFormat() string {
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	return nowTime
}

func TestTimeParse()  {
	t := time.Now()
	format := t.Format("2006-01-02 15:04:05")
	parse, err := time.Parse(defaultTime, format)
	if err!=nil {
		log.Println("error:",err)
	}
	log.Println("parse:",parse)
}

func main() {
	testTime := "2018-08-01 12:00:15"
	parse, _ := time.Parse(defaultTime, testTime)
	log.Println(parse)
	log.Println(time.Now().Format(defaultTime))
	TestTimeParse()

	// 获取当前时间戳
	log.Println(time.Now().Unix())

}