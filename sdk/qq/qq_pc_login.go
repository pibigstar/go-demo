package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	AppId       = "101827468"
	AppKey      = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	redirectURI = "http://127.0.0.1:9090/qqLogin"
)

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

func main() {
	http.HandleFunc("/toLogin", GetAuthCode)
	http.HandleFunc("/qqLogin", GetToken)

	fmt.Println("started...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}

// 1. Get Authorization Code
func GetAuthCode(w http.ResponseWriter, r *http.Request) {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	params.Add("state", "test")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)

	http.Redirect(w, r, loginURL, http.StatusFound)
}

// 2. Get Access Token
func GetToken(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("code", code)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/token", str)

	response, err := http.Get(loginURL)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)

	resultMap := convertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	GetOpenId(info, w)
}

// 3. Get OpenId
func GetOpenId(info *PrivateInfo, w http.ResponseWriter) {
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", "https://graph.qq.com/oauth2.0/me", info.AccessToken))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	body := string(bs)
	info.OpenId = body[45:77]

	GetUserInfo(info, w)
}

// 4. Get User info
func GetUserInfo(info *PrivateInfo, w http.ResponseWriter) {
	params := url.Values{}
	params.Add("access_token", info.AccessToken)
	params.Add("openid", info.OpenId)
	params.Add("oauth_consumer_key", AppId)

	uri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?%s", params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	w.Write(bs)
}

func convertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}
