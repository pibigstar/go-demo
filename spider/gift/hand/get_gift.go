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
	Name string `json:"name"`
}

// 尝试次数
const TryTime = 3
//手动
func main() {
	now := time.Now()
	date := now.Format("2006-01-02")
	for i := 0; i < TryTime; i++ {
		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://pay.qun.qq.com/cgi-bin/group_pay/good_feeds/draw_lucky_gift", strings.NewReader("bkn=183506344&from=0&gc=40636692&client=1&version=7.7.0.3645"))
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		req.Header.Set("Cookie", getCookie())

		resp, err := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s 请求内容失败,第%d次尝试 \n", date, i+1)
			continue
		}
		defer resp.Body.Close()
		var gift = Gift{}
		err = json.Unmarshal(body, &gift)
		if err != nil {
			fmt.Printf("%s JSON解析失败,第%d次尝试 \n", date, i+1)
			continue
		}
		if gift.Name == "" || len(gift.Name) == 0 {
			fmt.Printf("%s Cookie已失效,第%d次尝试 \n", date, i+1)
			break
		}
		fmt.Println(date + " 获得礼物:" + gift.Name)
		break
	}
}

func getCookie() string {
	bytes, err := ioutil.ReadFile("http/gift/cookie")
	if err != nil {
		fmt.Println("读取Cookie失败:", err.Error())
	}
	return string(bytes)
}
