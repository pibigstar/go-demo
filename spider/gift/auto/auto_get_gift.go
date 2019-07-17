package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Gift struct {
	Name string `json:"name"`
}

//自动获取
func main() {
	c := cron.New()
	// 每天凌晨5点执行一次
	spec := "0 0 5 * * ?"
	c.AddFunc(spec, func() {
		client := &http.Client{}

		req, err := http.NewRequest("POST", "https://pay.qun.qq.com/cgi-bin/group_pay/good_feeds/draw_lucky_gift", strings.NewReader("bkn=183506344&from=0&gc=40636692&client=1&version=7.7.0.3645"))
		if err != nil {
			fmt.Println("构建request失败")
		}

		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		req.Header.Set("Cookie", getCookie2())

		resp, err := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("请求内容失败")
		}
		defer resp.Body.Close()
		var gift = Gift{}
		err = json.Unmarshal(body, &gift)
		if err != nil {
			fmt.Println("JSON解析失败", err.Error())
		}
		now := time.Now()
		date := now.Format("2006-01-02")
		fmt.Println(date + " 获得礼物:" + gift.Name)
	})
	c.Start()

	select {}
}

func getCookie2() string {
	bytes, err := ioutil.ReadFile("cookie")
	if err != nil {
		fmt.Println("读取Cookie失败:", err.Error())
	}
	return string(bytes)
}
