package qq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetQQInfo(t QQType) (*User, error) {
	var user = &User{}
	user.Title = t.Title()
	client := http.Client{}
	// 1. 获取pt_local_token
	req, _ := http.NewRequest("GET", "https://xui.ptlogin2.qq.com/cgi-bin/xlogin?s_url="+t.TargetURL()+"&style=20&appid=715021417&proxy_url=https%3A%2F%2Fhuifu.qq.com%2Fproxy.html", nil)
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get pt_local_token, err: %s", err.Error())
	}
	ptLocalToken := processStr(response.Header["Set-Cookie"], "pt_local_token")
	user.PtLocalToken = ptLocalToken

	// 2.获取本机所登陆的QQ号码
	flag := false
	for i := 0; i < 8; i++ {
		req, _ = http.NewRequest("GET", fmt.Sprintf("https://localhost.ptlogin2.qq.com:430%d/pt_get_uins?callback=ptui_getuins_CB&r=0.6694805047494219&pt_local_tk=%s", i, ptLocalToken), nil)
		req.Header.Set("cookie", fmt.Sprintf("pt_local_token=%s", ptLocalToken))
		req.Header.Set("referer", t.Referer())
		res, err := client.Do(req)
		if err != nil || res == nil {
			continue
		}

		bytes, _ := ioutil.ReadAll(res.Body)
		body := string(bytes)
		r := regexp.MustCompile("\\[.*?]")
		temp := string(r.Find([]byte(body)))
		temp = temp[1 : len(temp)-1]
		json.Unmarshal([]byte(temp), &user)
		flag = true
		break
	}
	if !flag {
		return nil, fmt.Errorf("get localhost qq failed")
	}

	// 3. 获取clientkey
	req, _ = http.NewRequest("GET", fmt.Sprintf("https://localhost.ptlogin2.qq.com:4301/pt_get_st?clientuin=%s&callback=ptui_getst_CB&r=0.7284667321181328&pt_local_tk=%s", user.Account, ptLocalToken), nil)
	req.Header.Set("cookie", fmt.Sprintf("pt_local_token=%s", ptLocalToken))
	req.Header.Set("referer", t.Referer())
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get client key, err: %s", err.Error())
	}
	clientKey := processStr(res.Header["Set-Cookie"], "clientkey")

	// 4. 获取skey
	url := "https://ptlogin2.qq.com/jump?clientuin=" + user.Account + "&keyindex=9&pt_aid=549000912&daid=5&u1=" + t.TargetURL() + "&pt_local_tk=" + ptLocalToken + "&pt_3rd_aid=0&ptopt=1&style=40&has_onekey=1"
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", fmt.Sprintf("pt_local_token=%s;clientuin=%s;clientkey=%s", ptLocalToken, user.Account, clientKey))
	req.Header.Set("referer", t.Referer())

	res, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get skey, err: %s", err.Error())
	}

	// 获取uin和skey
	uin := processStr(res.Header["Set-Cookie"], "uin")
	skey := processStr(res.Header["Set-Cookie"], "skey")
	user.Uin = uin
	user.Skey = skey

	// 获取返回的URL
	all, _ := ioutil.ReadAll(res.Body)
	temp := string(all)
	r := regexp.MustCompile("https(.*?)'")
	temp = string(r.Find([]byte(temp)))
	url = temp[0 : len(temp)-1]

	// 5. 根据第四步返回的URL，获取p_skey
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", fmt.Sprintf("pt_local_token=%s", ptLocalToken))
	req.Header.Set("referer", t.Referer())
	res, err = client.Do(req)
	if err != nil {
		return user, fmt.Errorf("get uin and p_skey, err: %s", err.Error())
	}
	pSkey := processStr(res.Request.Response.Header["Set-Cookie"], "p_skey")
	user.PSkey = pSkey
	user.GTK = genderGTK(skey)

	return user, nil
}

// 根据key匹配数组中的值
func processStr(maps []string, key string) string {
	keyTemp := key + "="
	for _, v := range maps {
		if strings.Contains(v, keyTemp) && strings.Index(v, key) < 3 {
			r := regexp.MustCompile(keyTemp + "(.*?);")
			temp := string(r.Find([]byte(v)))
			temp = strings.Replace(temp, keyTemp, "", 1)
			value := temp[0 : len(temp)-1]
			return value
		}
	}
	return ""
}

// 根据skey计算出g_tk/bkn
func genderGTK(skey string) string {
	hash := 5381
	for _, s := range skey {
		us, _ := strconv.Atoi(fmt.Sprintf("%d", s))
		hash += (hash << 5) + us
	}
	return fmt.Sprintf("%d", hash&0x7fffffff)
}
