package proxy

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	urlpkg "net/url"
	"sort"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mnhkahn/maodou/request/goreq"
)

type XiciConfig struct {
	MaxCawlCnt int    `json:"max_cawl_cnt"`
	Cnt        int    `json:"cnt"`
	MinCnt     int    `json:"min_cnt"`
	Root       string `json:"root"`
}

type XiciProxy struct {
}

func (this *XiciProxy) NewProxyImpl(dsn string) (ProxyContainer, error) {
	d := new(XiciProxyContainer)
	config := new(XiciConfig)
	err := json.Unmarshal([]byte(dsn), config)
	d.config = config
	if err != nil {
		return d, fmt.Errorf("Config for xici is error: %v", err)
	}

	return d, nil
}

type XiciProxyContainer struct {
	config  *XiciConfig
	req     goreq.Request
	proxies ProxyConfigs
}

func (this *XiciProxyContainer) Init() {
	this.init()
}

func (this *XiciProxyContainer) init() {
	resp, err := goreq.Request{Uri: "http://www.xicidaili.com/"}.Do()
	if err == nil {
		raw_document, _ := html.Parse(resp.Body)
		document := goquery.NewDocumentFromNode(raw_document)
		document.Find(`#ip_list > tbody > tr`).Each(func(i int, s *goquery.Selection) {
			if i > 1 && len(this.proxies) < this.config.Cnt && s.Children().Length() == 7 {
				p := new(ProxyConfig)
				p.Ip = s.Children().Get(1).FirstChild.Data
				p.Port, _ = strconv.Atoi(s.Children().Get(2).FirstChild.Data)
				if s.Children().Get(3).FirstChild != nil {
					p.Location = s.Children().Get(3).FirstChild.Data
				}
				p.Safe = s.Children().Get(4).FirstChild.Data == "高匿"
				p.Schema = s.Children().Get(5).FirstChild.Data
				p.VerifyTime = s.Children().Get(6).FirstChild.Data
				if this.TestProxy(p) {
					p.Id = len(this.proxies)
					this.add(p)
				}
			}
		})
	} else {
		log.Println(err)
	}
}

func (this *XiciProxyContainer) add(p *ProxyConfig) {
	log.Printf("Got proxy %v.", p)
	exist_flag := false
	for _, temp := range this.proxies {
		if temp.Ip == p.Ip {
			exist_flag = true
			break
		}
	}
	if !exist_flag {
		this.proxies = append(this.proxies, p)
	}
}

func (this *XiciProxyContainer) One() *ProxyConfig {
	if len(this.proxies) == 0 {
		return new(ProxyConfig)
	}

	sort.Sort(this.proxies)
	p := this.proxies[0]
	p.Cnt++

	if p.Cnt >= this.config.MaxCawlCnt {
		this.DeleteProxy(0)
	}
	return p
}

func (this *XiciProxyContainer) Len() int {
	return len(this.proxies)
}

func (this *XiciProxyContainer) Update(p *ProxyConfig) {
	p.Delayed = p.Delayed / time.Duration(p.Cnt+1)
	this.add(p)
}

func (this *XiciProxyContainer) DeleteProxy(i int) {
	for idx, temp := range this.proxies {
		if temp.Id == i {
			log.Printf("Delete proxy %v", temp)
			this.proxies = append(this.proxies[:idx], this.proxies[idx+1:]...)
			break
		}
	}

	if len(this.proxies) <= this.config.MinCnt {
		this.Init()
	}
}

// true means proxy is OK
func (this *XiciProxyContainer) TestProxy(p *ProxyConfig) bool {
	start := time.Now()
	if p.Ip == "" {
		return false
	}
	u := new(urlpkg.URL)
	u.Scheme = "http"
	u.Host = fmt.Sprintf("%s:%d", p.Ip, p.Port)
	res, err := goreq.Request{Uri: this.config.Root, Proxy: u.String(), Timeout: time.Duration(5) * time.Second}.Do()
	p.Delayed = time.Now().Sub(start)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	if err != nil || res == nil || res.StatusCode != http.StatusOK {
		log.Println("Proxy test failed", p)
		return false
	}
	return true
}

func init() {
	Register("xici", &XiciProxy{})
}
