package weixin

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	// 统一下单API接口
	WeiXinPayUnifedURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	// 通知地址
	WeixinNotifyURL = "https://pibigstar.com/weixin/pay"
	// 交易类型
	TradeTypeJSAPI  = "JSAPI"
	TradeTypeNATIVE = "NATIVE"
	// 商品描述
	RewardBody = "微信支付Demo"
	// 终端IP，用户的客户端IP
	CreateIP = "127.0.0.1"
	// 微信支付分配的公众账号ID（企业号corpid即为此appId）
	AppID = "wx8888888888888888"
	// 商户号
	MchID = "1230000109"
	// 商户Key
	MchKey = "pibigstar"
)

type PayRequest struct {
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
	SignType  string `json:"signType"`
}

// 微信支付
func WeiXinPay(amount int32, openId string) (*PayRequest, error) {
	// 商户server调用统一下单接口请求订单
	response, err := createUnifyOrder(amount, openId, TradeTypeJSAPI)
	if err != nil {
		return nil, err
	}

	timeStamp := fmt.Sprintf("%d", time.Now().Unix())
	nonceStr := genNonceStr(16)
	// 预支付交易会话标识
	// 用于后续接口调用中使用，该值有效期为2小时
	packageStr := "prepay_id=" + response.PrepayID

	// 生成支付签名
	paySign := buildSign(map[string]string{
		"appId":     AppID,
		"timeStamp": timeStamp,
		"nonceStr":  nonceStr,
		"package":   packageStr,
		"signType":  "MD5",
	})

	// 使用：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
	return &PayRequest{
		TimeStamp: timeStamp,
		NonceStr:  nonceStr,
		Package:   packageStr,
		PaySign:   paySign,
		SignType:  "MD5",
	}, nil
}

// 微信扫码支付
// 生成支付二维码
func WeixinPayCode(amount int32, openId string) (string, error) {
	// 商户server调用统一下单接口请求订单
	response, err := createUnifyOrder(amount, openId, TradeTypeNATIVE)
	if err != nil {
		return "", err
	}
	// 当trade_type=NATIVE时才有返回
	codeURL := response.CodeURL
	fmt.Println(codeURL)

	return codeURL, nil
}

// 统一下单
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func createUnifyOrder(amount int32, openId, tradeType string) (*UnifyOrderResponse, error) {
	bodyData := buildUnifyParams(amount, openId, tradeType)
	requestParams, err := xml.Marshal(bodyData)
	if err != nil {
		return nil, err
	}

	strReq := string(requestParams)
	strReq = strings.Replace(strReq, "UnifyOrderRequest", "xml", -1)
	response, err := http.Post(WeiXinPayUnifedURL, "text/xml:charset=UTF-8", strings.NewReader(strReq))
	if err != nil {
		return nil, err
	}

	response.Header.Set("Content-Type", "text/xml;charset=utf-8")
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	uoResp := new(UnifyOrderResponse)
	if _err := xml.Unmarshal(responseBody, uoResp); _err != nil {
		return nil, _err
	}

	createOrder(bodyData.OutTradeNo, amount)
	return uoResp, nil
}
func createOrder(outTradeNo string, amount int32) {
	// 将此订单存储到数据库中
	fmt.Printf("save to db, outTradeNo: %s, amount: %d", outTradeNo, amount)
}

// 统一下单请求体
type UnifyOrderRequest struct {
	AppID          string `xml:"appid"`
	MchID          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Body           string `xml:"body"`
	OutTradeNo     string `xml:"out_trade_no"`
	TotalFee       string `xml:"total_fee"`
	SpbillCreateIP string `xml:"spbill_create_ip"`
	NotifyURL      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
	OpenID         string `xml:"openid"`
}

// 统一下单响应体
type UnifyOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	ReturnMsg  string `xml:"return_msg"`
	PrepayID   string `xml:"prepay_id"`
	CodeURL    string `xml:"code_url"`
	NoneceStr  string `xml:"nonce_str"`
}

// 构造统一下单参数
func buildUnifyParams(amount int32, openID, tradeType string) *UnifyOrderRequest {
	unifyParams := UnifyOrderRequest{
		AppID:          AppID,
		MchID:          MchID,
		NonceStr:       genNonceStr(16),
		Body:           RewardBody,
		OutTradeNo:     buildOutTradeNo(),
		TotalFee:       fmt.Sprintf("%d", amount),
		SpbillCreateIP: CreateIP,
		NotifyURL:      WeixinNotifyURL,
		TradeType:      tradeType,
		OpenID:         openID,
	}

	sign := buildSign(map[string]string{
		"appid":            unifyParams.AppID,
		"mch_id":           unifyParams.MchID,
		"nonce_str":        unifyParams.NonceStr,
		"body":             unifyParams.Body,
		"out_trade_no":     unifyParams.OutTradeNo,
		"total_fee":        unifyParams.TotalFee,
		"spbill_create_ip": unifyParams.SpbillCreateIP,
		"notify_url":       unifyParams.NotifyURL,
		"trade_type":       unifyParams.TradeType,
		"openid":           unifyParams.OpenID,
	})

	unifyParams.Sign = sign
	return &unifyParams
}

// 获取指定长度随机字符串
func genNonceStr(n int) string {
	letterBytes := "1234567890abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}

// 生成商户订单号
func buildOutTradeNo() string {
	timeStamp := time.Now().UnixNano()
	outTradeNoStr := fmt.Sprintf("%d", timeStamp)
	h := md5.New()
	io.WriteString(h, outTradeNoStr)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 生成签名
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
func buildSign(params map[string]string) string {
	var signString string
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if params[key] != "" {
			signString += key + "=" + params[key] + "&"
		}
	}

	signString += "key=" + MchKey

	h := md5.New()
	io.WriteString(h, signString)
	signString = strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	return signString
}
