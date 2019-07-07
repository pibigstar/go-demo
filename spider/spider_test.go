package spider

import (
	"encoding/xml"
	"fmt"
	"go-demo/spider/agent"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

//chan中存入string类型的href属性,缓冲200
var urlChannel = make(chan string, 200)

//以Must前缀的方法或函数都是必须保证一定能执行成功的,否则将引发一次panic
var atagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`)

func TestSpider(t *testing.T) {
	go Spider("http:/blog.csdn.net")
	//for url := range urlChannel {
	//	//通过runtime可以获取当前运行时的一些相关参数等
	//	fmt.Println("routines num = ", runtime.NumGoroutine(), "chan len = ", len(urlChannel))
	//	go Spider(url)
	//}
}

func Spider(url string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", agent.GetRandomUserAgent())
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		fmt.Errorf("Get请求%s返回错误:%s", url, e)
		return
	}

	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		atag := atagRegExp.FindAllString(resStr, -1)
		for _, a := range atag {
			href, _ := GetHref(a)
			if strings.Contains(href, "article/details/") {
				fmt.Println("☆", href)
			} else {
				fmt.Println("□", href)
			}
			urlChannel <- href
		}
	}
}

func GetHref(atag string) (href, content string) {
	inputReader := strings.NewReader(atag)
	decoder := xml.NewDecoder(inputReader)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				if strings.EqualFold(attrName, "href") || strings.EqualFold(attrName, "HREF") {
					href = attrValue
				}
			}
			// 处理元素结束（标签）
		case xml.EndElement:
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content = string([]byte(token))
		default:
			href = ""
			content = ""
		}
	}
	return href, content
}
