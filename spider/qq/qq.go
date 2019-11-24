package qq

type User struct {
	Title        string `json:"title"`
	Account      string `json:"account"`
	Nickname     string `json:"nickname"`
	PtLocalToken string `json:"pt_local_token"`
	Uin          string `json:"uin"`
	Skey         string `json:"skey"`
	PSkey        string `json:"p_skey"`
	GTK          string `json:"g_tk"`
}

const (
	// qq空间
	qzoneReferer   = "https://xui.ptlogin2.qq.com/cgi-bin/xlogin?proxy_url=https%3A//qzs.qq.com/qzone/v6/portal/proxy.html&daid=5&&hide_title_bar=1&low_login=0&qlogin_auto_login=1&no_verifyimg=1&link_target=blank&appid=549000912&style=22&target=self&s_url=https%3A%2F%2Fqzs.qzone.qq.com%2Fqzone%2Fv5%2Floginsucc.html%3Fpara%3Dizone&pt_qr_app=%E6%89%8B%E6%9C%BAQQ%E7%A9%BA%E9%97%B4&pt_qr_link=http%3A//z.qzone.com/download.html&self_regurl=https%3A//qzs.qq.com/qzone/v6/reg/index.html&pt_qr_help_link=http%3A//z.qzone.com/download.html&pt_no_auth=1"
	qzoneTargetURL = "https://qzs.qzone.qq.com/qzone/v5/loginsucc.html"

	// qq好友
	friendReferer   = "https://xui.ptlogin2.qq.com/cgi-bin/xlogin?pt_disable_pwd=1&appid=715030901&daid=73&pt_no_auth=1&s_url=https%3A%2F%2Fqun.qq.com%2Fmanage.html"
	friendTargetURL = "https://qun.qq.com/member.html"
)

type QQType int

const (
	QQZone   QQType = 1
	QQFriend QQType = 2
)

func (t QQType) Referer() string {
	switch t {
	case QQZone:
		return qzoneReferer
	case QQFriend:
		return friendReferer
	default:
		return ""
	}
}

func (t QQType) Title() string {
	switch t {
	case QQZone:
		return "QZone"
	case QQFriend:
		return "Friends"
	default:
		return ""
	}
}

func (t QQType) TargetURL() string {
	switch t {
	case QQZone:
		return qzoneTargetURL
	case QQFriend:
		return friendTargetURL
	default:
		return ""
	}
}
