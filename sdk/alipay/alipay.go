package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/smartwalle/alipay/v3"
)

var (
	appId = "2016091800540000"
	//生成CSR文件中的私钥
	privateKey = "MIIEowIBAAKCAQEA5hhBBH4CSO1AySXprGazNZKVR/+G87GHt1U2SQL33WCFPKFYqE0qlpdLKECLb6xgOlSGFD676hxJcBH//l+A16Wb0WjVAbIIc7ichjantyg9t3FQP97t9wI5mYlUGMsTu2TqzFu/0hq/KAhSeNlg15m3v+4g81qITnTapzFAbC/6LLw9Sznq+eKi3bK3scAWZtkpJRyT/UqEXDsixskMzezISqYOQW9xHdfBvkzcGKuxi/N0z6rfoxyG2D1fNQ6dQ91vb54Ph0rdojDKUB29CdY92KoxPvAumXBgVf6k3MRpF4di0c9lR5cqgpbTUkotcuidCWgx8w13HNenpF2dKQIDAQABAoIBAASZa4NJeYY3p9nddiRKET765R0BUJNCczII9ALVmlrEeSVTHFCQ6k8ESy5My/y5d1rzIZL6BguR8S3aTkGpawvkdY7kB433HxAhGo/cO9H/bexiyXXdYOhVFQ2qnxG3zXcrdz4Kf3UVr8h/Ehb0UWk921xsyB/VKXBYCZ7Z7y26ZhO6J+ZRHqGgAW7vXeXmhhAyIixIuUQnVkRSXrKKnx+HHYxhioIgKUtny9N/636jnC9rMi2mKK5cjgPjgX4EVJMNTn/jdTNrYZxgg6F+innwvyHpMGggxiCMpTjtC3efujOocMmMgNSb1J7NVdqOW22n2MNlG5JLgAktLs4HYnkCgYEA9X15fkrQYqahNuvgUZtGofZQHPwXXOLcP/t4rZcdunwKuyjMaiFnhBX3p8WBaVi0zwIXe86OkaLWZhalbDpmcGnZRaUVJq9KLnIaA9QDectnoqDpiR1IxFfKM2IsIYRU8Y5+OKQwXtAHfuEbySnb++McTUPUB9poZI6BswxWngsCgYEA7/IMfgejlRPQpysecyDx1y5x6ndnZ2VapJdaS5Q/g+1j/ZjX52oo3qeEA3DVBmDGQnvo8voCcdtLIp+vITEEsEtiDWrJk3giGF73HaPvRqrYtH5uhLLQy1Ehub0Nc/jWSrV3S2bSir6Qmk9ptu4ndGSOeiFziS/iwrFF70ayFhsCgYBGfDVjBpYYjSFixI0OwVehbziHafZHTDfTAyAeL3Jwteba4Bb5Lggry6bk+/dxSO/5M++MM72JoUiP3Va34XjCNBIXRhPxnIjfFxHTIY+x664g6rTDEq5u+YnsAPcM1JMTHEevea0NvAs66eVxd9xa0VWx9ZSugI5SuPwSbat9CwKBgQCaVNx2H6G23GTjcReHw5Pp7OS2g5CN76IKpZMdc8AashETZ0DPhve8ppCBygwqqwo6bwqZZfc2lm9QWNdDCQ1T+1iY+qum36lGdaaKeQwJLxBtn7ikP4OOkqOXnSLPCimDKg8N/5fCR+ooZpW/ZJUaByehJGz0u0kmIvGxgo4/KwKBgAYbhMiE0IMsAYr5dPhX52FWaSRzcF/5WALFIOes5jCctSrMFMVPMq0xp1QWFK6iHeA1nv8Uk2NOQXbn13iTr+LTD01aT0WUJ7r5fhatrvfqIUHm7tcTzgwjui6QYv0y2m3hUXA1wCwMW9OL4Eadk32x4okssyMxWbJdCRvX4Ssz"
	client     *alipay.Client
)

func init() {
	client, _ = alipay.New(appId, privateKey, false)
	client.LoadAppPublicCertFromFile("appPublic.crt")
	client.LoadAliPayPublicCertFromFile("aliPayPublic.crt")
	client.LoadAliPayRootCertFromFile("aliPayRoot.crt")
}

// 调转支付宝网站支付
func WebPageAlipay() {
	pay := alipay.TradePagePay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://www.pibigstar/alipay"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8088/return"
	//支付标题
	pay.Subject = "网页扫码支付"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = "0.01"
	v := url.Values{}
	v.Add("orgId", "123456")
	body := v.Encode()
	// 支付宝回传参数，需要经过UrlEncode
	pay.PassbackParams = body

	url, err := client.TradePagePay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	//打开默认浏览器
	payURL = strings.Replace(payURL, "&", "^&", -1)
	exec.Command("cmd", "/c", "start", payURL).Start()
}

// https://docs.open.alipay.com/204/105695/
// 手机网页支付(可转到APP支付）
func WapAlipay() {
	pay := alipay.TradeWapPay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://www.pibigstar/alipay"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8088/return"
	//支付标题
	pay.Subject = "手机网页支付"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()
	//商品code
	pay.ProductCode = "QUICK_WAP_WAY"
	//支付失败后返回地址
	pay.QuitURL = "http://www.pibigstar/alipay"
	//金额
	pay.TotalAmount = "0.01"
	v := url.Values{}
	v.Add("orgId", "123456")
	body := v.Encode()
	// 支付宝回传参数，需要经过UrlEncode
	pay.PassbackParams = body

	url, err := client.TradeWapPay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)
	//打开默认浏览器
	payURL = strings.Replace(payURL, "&", "^&", -1)
	exec.Command("cmd", "/c", "start", payURL).Start()
}

// APP支付
func GetAppPayURL() {
	pay := alipay.TradeAppPay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://www.pibigstar/alipay"
	//支付标题
	pay.Subject = "APP支付"

	//二维码使用的Code
	pay.ProductCode = time.Now().String()
	pay.TotalAmount = "10.00"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()

	result, err := client.TradeAppPay(pay)
	if err != nil {
		fmt.Println(err)
	}
	// 跳转APP支付参数
	fmt.Println(result)
}

// https://docs.open.alipay.com/api_1/alipay.trade.pay/
// 扫码支付(生成支付的二维码的链接）
func GetQrPayURL() {
	pay := alipay.TradePreCreate{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://www.pibigstar/alipay"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8088/return"
	//支付标题
	pay.Subject = "支付宝支付测试"

	//二维码使用的Code
	pay.ProductCode = "FACE_TO_FACE_PAYMENT"
	pay.TotalAmount = "10.00"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()

	url, err := client.TradePreCreate(pay)
	if err != nil {
		fmt.Println(err)
	}
	//二维码链接，可用此链接生成一个二维码扫码支付
	fmt.Println(url.Content.QRCode)
}

func main() {
	//生成支付URL
	WapAlipay()
	//支付成功之后的返回URL页面
	http.HandleFunc("/return", func(rep http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ok, err := client.VerifySign(req.Form)
		if err == nil && ok {
			rep.Write([]byte("支付成功"))
		}
	})
	//支付成功之后的通知页面
	http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
		var notification, err = client.GetTradeNotification(req)
		if err != nil {
			fmt.Println("支付失败")
			rep.WriteHeader(http.StatusForbidden)
			return
		}
		values, err := url.ParseQuery(notification.PassbackParams)
		if err != nil {
			fmt.Println("解析参数失败")
			rep.WriteHeader(http.StatusForbidden)
			return
		}
		fmt.Println(values.Get("orgId"))
		//支付宝支付成功之后的信息
		fmt.Printf("%+v", notification)
		fmt.Println("支付成功")
		//修改订单状态。。。。
		alipay.AckNotification(rep) // 确认收到通知消息
	})

	fmt.Println("server start....")
	http.ListenAndServe(":8088", nil)
}
