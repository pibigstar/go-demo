Go Weibo SDK
========
Sina Weibo SDK For Gopher

[![Build Status](https://drone.io/github.com/going/weigo/status.png)](https://drone.io/github.com/going/weigo/latest)

文档请看测试用例，哈哈！

WIP 部分高级功能啥的还没有添加静态方法。。常用的都已经添加了。。有需求请提啊！

##Install:
```go
go get -u github.com/xiocode/weigo
```

##Usage:
http://open.weibo.com/wiki/API%E6%96%87%E6%A1%A3_V2
参照官方文档调用对应的方法.目前weigo支持的功能可以参考本项目的[wiki](https://github.com/violetgo/weigo/wiki).
```go
package weigo

import (
	"testing"
)

var api *APIClient

func init() {
	if api == nil {
		api = NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
		api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	}
}

// 先获取 GetAuthorizeUrl 打开浏览器操作之后 得到 http://127.0.0.1/callback?code=198ea555a7efbbfd90caa92c86feb2b5
// code的值是RequestAccessToken的参数
func TestGetAuthorizeUrl(t *testing.T) {
	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
	authorize_url, err := api.GetAuthorizeUrl(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authorize_url)
}

func TestRequestAccessToken(t *testing.T) {
	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
	var result map[string]interface{}
	err := api.RequestAccessToken("1fdaa295b73d2a9568e284383ced5e9e", &result) // code
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	access_token := result["access_token"]
	fmt.Println(reflect.TypeOf(access_token), access_token)
	expires_in := result["expires_in"]
	fmt.Println(reflect.TypeOf(expires_in), expires_in)
}

func Test_GET_statuses_user_timeline(t *testing.T) {
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(Statuses)
	err := api.GET_statuses_user_timeline(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Statuses))
}

func Test_GET_statuses_home_timeline(t *testing.T) {
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(Statuses)
	err := api.GET_statuses_home_timeline(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Statuses))
}

func Test_GET_statuses_repost_timeline(t *testing.T) {
	kws := map[string]interface{}{
		"id": "3551749023600582",
	}
	result := new(Reposts)
	err := api.GET_statuses_repost_timeline(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Reposts))
}

func Test_POST_statuses_repost(t *testing.T) {
	kws := map[string]interface{}{
		"id": "3551749023600582",
	}
	result := new(Status)
	err := api.POST_statuses_repost(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}

func Test_POST_statuses_repost(t *testing.T) {
	kws := map[string]interface{}{
		"status": "Testing...Testing...",
	}
	result := new(Status)
	err := api.POST_statuses_update(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}

func Test_POST_statuses_repost(t *testing.T) {
	kws := map[string]interface{}{
		"id": "3556138715301190",
	}
	result := new(Status)
	err := api.POST_statuses_destroy(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}

```

Weibo: http://weibo.com/xceman @XIOCODE
Gmail: xiocode@gmail.com
