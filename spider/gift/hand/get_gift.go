package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Gift struct {
	Name   string `json:"name"`
	LuckId int    `json:"luckid"`
}

const (
	// 尝试次数
	TryTime = 3
	// 请求地址
	URL = "https://pay.qun.qq.com/cgi-bin/group_pay/good_feeds/draw_lucky_gift"
)

//手动
func main() {
	now := time.Now()
	date := now.Format("2006-01-02")
	for i := 0; i < TryTime; i++ {
		client := &http.Client{}
		req, err := http.NewRequest("POST", URL, strings.NewReader("bkn=1858570791&from=0&gc=594174090&client=1&version=0"))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		req.Header.Set("Cookie", getCookie())

		resp, err := client.Do(req)
		if err != nil {
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s 请求内容失败,第%d次尝试 \n", date, i+1)
			continue
		}
		defer resp.Body.Close()

		if strings.Contains(string(body), "basekey") {
			fmt.Printf("%s Cookie已失效,第%d次尝试 \n", date, i+1)
			break
		}

		var gift Gift
		err = json.Unmarshal(body, &gift)
		if err != nil {
			fmt.Printf("%s JSON解析失败,第%d次尝试 \n", date, i+1)
			continue
		}
		if gift.Name == "" || gift.LuckId == 0 {
			fmt.Println(date + " 未获得礼物")
			break
		}
		fmt.Println(date + " 获得礼物:" + gift.Name)
		break
	}
}

func getCookie() string {
	bytes, err := ioutil.ReadFile("spider/gift/cookie")
	if err != nil {
		fmt.Println("读取Cookie失败:", err.Error())
	}
	return string(bytes)
}
