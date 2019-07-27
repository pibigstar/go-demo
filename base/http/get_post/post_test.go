package get_post

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestPostForm(t *testing.T) {
	response, err := http.PostForm(
		"http://www.baidu.com",
		url.Values{
			"username": {"rob"},
			"password": {"abc123_"},
		})
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

func TestHttpPost(t *testing.T) {
	// contentType必须是这个，不然会报错
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
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

	req, err := http.NewRequest(
		"POST",
		"http://www.baidu.com",
		strings.NewReader("name=pibigstar"),
	)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}
