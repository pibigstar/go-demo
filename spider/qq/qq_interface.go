package qq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAllFriends(user *User) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://qun.qq.com/cgi-bin/qun_mgr/get_friend_list", strings.NewReader(fmt.Sprintf("bkn=%s", user.GTK)))
	req.Header.Set("cookie", fmt.Sprintf("uin=%s;p_uin=%s;skey=%s;p_skey=%s", user.Uin, user.Uin, user.Skey, user.PSkey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("referer", "https://qun.qq.com/member.html")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
