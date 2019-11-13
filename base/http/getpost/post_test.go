package getpost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// 提交Form表单数据
func TestPostForm(t *testing.T) {
	params := url.Values{}
	params.Add("name", "pibigstar")
	params.Add("age", "18")

	// Content-Type = application/x-www-form-urlencoded
	response, err := http.PostForm("http://www.baidu.com", params)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

// 提交JSON数据
func TestHttpPostJson(t *testing.T) {
	user := &User{
		UserName: "pibigstar",
		Password: "123456",
	}
	bs, _ := json.Marshal(user)
	data := bytes.NewReader(bs)

	resp, err := http.Post("http://www.baidu.com", "application/json", data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}

// 有复杂请求，设置Header，cookie等需要使用这个
func TestHttpDo(t *testing.T) {
	client := &http.Client{}

	params := url.Values{}
	params.Add("name", "pibigstar")
	params.Add("age", "18")

	m := make(map[string]string)
	for key := range params {
		m[key] = params.Get(key)
	}
	bs, _ := json.Marshal(m)
	data := bytes.NewReader(bs)

	req, err := http.NewRequest("POST", "http://www.baidu.com", data)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	req.Header.Set("Cookie", "name=anny")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
