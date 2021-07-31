package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"io/ioutil"
	logs "log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	communicatedNo  = 1 // 未被同事沟通过
	communicatedYes = 2 // 已被同事沟通过

	requestTypeToMe   = "3" // 牛人向我发简历
	requestTypeToGeek = "4" // 我向牛人发求简历

	cookieFile    = "cookie.txt"
	school985File = "985.txt"
	school211File = "211.txt"
	jobsFile      = "jobs.txt"
	bossLog       = "boss.log"
)

var (
	jobIds = make(map[string]string) // 工作Id
	talked sync.Map                  // 已经沟通过的人

	maxLimit  = errors.New("今日沟通已达上限")
	notFriend = errors.New("好友关系校验失败")
	notLogin  = errors.New("当前登录状态已失效")

	runningTime = time.Minute * 3 // 只进行3分钟候选人选择
	school985   []string
	school211   []string
	cookie      string

	logFile, _ = os.OpenFile(bossLog, os.O_RDWR|os.O_CREATE, 0664)
	log        = logs.New(logFile, "", logs.Ldate|logs.Ltime)
)

func init() {
	// 读取cookie信息
	readCookie()
	// 监听cookie文件
	watchCookie()
	// 读取jobId
	readJobs()
	// 读取学校信息
	readSchool()
}

// 招人
func Hiring() {
	var wg sync.WaitGroup
	for jobId, jobName := range jobIds {
		fmt.Println("正在沟通职位:", jobName)
		wg.Add(1)
		go func(jobId string) {
			defer wg.Done()
			defer func() {
				if e := recover(); e != nil {
					log.Println("recover", e)
				}
			}()
			var (
				geeksQueue []*Geek
				ctx, _     = context.WithTimeout(context.Background(), runningTime)
				t          = time.NewTicker(5 * time.Second) // 5秒一次，防止被反爬
				jWg        sync.WaitGroup
			)
			for {
				select {
				case <-ctx.Done():
					// 开始打招呼
					if len(geeksQueue) > 0 {
						sort.Sort(SortGeek(geeksQueue))

						for _, l := range geeksQueue {
							if _, ok := talked.Load(l.GeekCard.GeekID); ok {
								continue
							}
							log.Printf("正在与: %s 打招呼 \n", l.GeekCard.GeekName)
							err := hello(jobId, l.GeekCard.EncryptGeekID, l.GeekCard.Lid, l.GeekCard.SecurityID, l.GeekCard.GeekID, l.GeekCard.ExpectID)
							if err == maxLimit {
								fmt.Println("今日已达上限")
								return
							}
							// 标记
							talked.Store(l.GeekCard.GeekID, "")

							// 轮询向牛人直接请求简历直到对方回复我们建立好友关系
							jWg.Add(1)
							go func(securityId string) {
								defer jWg.Done()
								t := time.NewTicker(time.Minute * 1)
								for {
									select {
									case <-t.C:
										if err := requestResumes(securityId); err == nil {
											t.Stop()
											return
										}
									}
								}
							}(l.GeekCard.SecurityID)
						}
						jWg.Wait()
						return
					}
				case <-t.C:
					geeks := searchGeekByJobId(jobId)
					geeksQueue = append(geeksQueue, geeks...)
				}
			}

		}(jobId)
	}
	wg.Wait()
}

func searchGeekByJobId(jobId string) []*Geek {
	var geeks []*Geek
	geekList, err := listRecommend(jobId)
	if err != nil {
		if err == notLogin {
			// todo: 发送邮件提醒
			panic(err)
		}
	}
	for _, geek := range geekList {
		log.Printf("候选人: %s  期待职位：%s \n", geek.GeekCard.GeekName, geek.PositionName)
		if selectGeek(geek) {
			fmt.Printf("候选人: %s  进入队列\n", geek.GeekCard.GeekName)
			geeks = append(geeks, geek)
		}
	}
	return geeks
}

// 筛选
// 985,211,大厂优先选择
func selectGeek(geek *Geek) bool {
	// 已经打过招呼了
	if geek.HaveChatted == 1 {
		return false
	}
	// 已经被同事撩过
	if geek.Cooperate == communicatedYes {
		return false
	}
	//  是否是本科, 2分
	if geek.GeekCard.GeekDegree == "本科" {
		geek.Weight += 2
	}
	// 是否是985
	if is985(geek.GeekCard.GeekEdu.School) {
		geek.Weight += 5
	}
	// 是否是211
	if is211(geek.GeekCard.GeekEdu.School) {
		geek.Weight += 3
	}
	// 在职-暂不考虑
	if strings.Contains(geek.GeekCard.ApplyStatusDesc, "暂不考虑") {
		geek.Weight += 1
	}
	// 在职-月内到岗
	if strings.Contains(geek.GeekCard.ApplyStatusDesc, "月内到岗") {
		geek.Weight += 2
	}
	// 离职-随时到岗
	if strings.Contains(geek.GeekCard.ApplyStatusDesc, "离职") {
		geek.Weight += 3
	}
	return true
}

func is985(school string) bool {
	for _, s := range school985 {
		if strings.EqualFold(s, school) {
			return true
		}
		if strings.Contains(school, s) {
			return true
		}
		if strings.Contains(s, school) {
			return true
		}
	}
	return false
}

func is211(school string) bool {
	for _, s := range school211 {
		if strings.EqualFold(s, school) {
			return true
		}
		if strings.Contains(school, s) {
			return true
		}
		if strings.Contains(s, school) {
			return true
		}
	}
	return false
}

// 打招呼
// 需要设置自动打招呼
func hello(jobId, encryptGeekId, lid, securityId string, geekId, expectId int) error {
	uri := fmt.Sprintf("https://www.zhipin.com/wapi/zpboss/h5/chat/start?_=%d", time.Now().Unix())
	urlQuery := url.Values{}
	urlQuery.Add("jid", jobId)
	urlQuery.Add("gid", encryptGeekId)
	urlQuery.Add("lid", lid)
	urlQuery.Add("expectId", fmt.Sprintf("%d", expectId))
	urlQuery.Add("securityId", securityId)

	data := strings.NewReader(urlQuery.Encode())
	req, _ := http.NewRequest(http.MethodPost, uri, data)
	addHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("hello request", err.Error())
		return err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	str := string(bs)
	if strings.Contains(str, "今日沟通已达上限") {
		return maxLimit
	}
	return nil
}

// 接收简历
func acceptResumes(mid, securityId string) error {
	uri := "https://www.zhipin.com/wapi/zpchat/exchange/accept"
	urlQuery := url.Values{}
	urlQuery.Add("mid", mid)
	urlQuery.Add("type", requestTypeToMe)
	urlQuery.Add("securityId", securityId)

	req, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(urlQuery.Encode()))
	addHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("acceptResumes request", err.Error())
		return err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
	return nil
}

// 向牛人请求简历
// 每隔一段时间请求一次，直到对方回复我们，建立好友关系为止
func requestResumes(securityId string) error {
	uri := "https://www.zhipin.com/wapi/zpchat/exchange/request"
	urlQuery := url.Values{}
	urlQuery.Add("type", requestTypeToGeek)
	urlQuery.Add("securityId", securityId)

	req, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(urlQuery.Encode()))
	addHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
	str := string(bs)
	if strings.Contains(str, "好友关系校验失败") {
		return notFriend
	}
	return nil
}

// 获取推荐牛人列表
func listRecommend(jobId string) ([]*Geek, error) {
	uri := fmt.Sprintf("https://www.zhipin.com/wapi/zprelation/interaction/bossGetGeek?")
	urlQueue := url.Values{}
	urlQueue.Add("gender", "0")
	urlQueue.Add("exchangeResumeWithColleague", "0")
	urlQueue.Add("switchJobFrequency", "0")
	urlQueue.Add("activation", "0")
	urlQueue.Add("recentNotView", "0")
	urlQueue.Add("school", "0")
	urlQueue.Add("major", "0")
	urlQueue.Add("experience", "0")
	urlQueue.Add("jobid", jobId)
	urlQueue.Add("degree", "0")
	urlQueue.Add("salary", "0")
	urlQueue.Add("intention", "0")
	urlQueue.Add("refresh", fmt.Sprintf("%d", time.Now().Unix()))
	urlQueue.Add("status", "1")
	urlQueue.Add("cityCode", "")
	urlQueue.Add("businessId", "0")
	urlQueue.Add("source", "")
	urlQueue.Add("districtCode", "0")
	urlQueue.Add("page", fmt.Sprintf("%d", 1))
	urlQueue.Add("tag", "1")

	uri = uri + urlQueue.Encode()
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	addHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ListRecommend request", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bs), "当前登录状态已失效") {
		return nil, notLogin
	}
	var temp *GeekListResp
	err = json.Unmarshal(bs, &temp)
	if err != nil {
		return nil, err
	}
	return temp.ZpData.GeekList, nil
}

func addHeader(req *http.Request) {
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", cookie)
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
}

// 监听cookie变化
func watchCookie() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	err = watcher.Add(cookieFile)
	if err != nil {
		log.Println("watch cookie.txt", err.Error())
		return
	}
	// 开始监听
	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				// cookie文件有变动，重新设置cookie
				readCookie()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watcher error:", err)
			}
		}
	}()
}

func readCookie() {
	bs, _ := ioutil.ReadFile(cookieFile)
	cookie = string(bs)
}

func readSchool() {
	bs, _ := ioutil.ReadFile(school985File)
	br := bufio.NewReader(bytes.NewReader(bs))
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		school985 = append(school985, string(a))
	}

	bs, _ = ioutil.ReadFile(school211File)
	br = bufio.NewReader(bytes.NewReader(bs))
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		school211 = append(school211, string(a))
	}
}

func readJobs() {
	bs, _ := ioutil.ReadFile(jobsFile)
	br := bufio.NewReader(bytes.NewReader(bs))
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := string(a)
		var (
			jobId   string
			jobName string
		)
		if ss := strings.Split(s, "//"); len(s) > 1 {
			jobId = strings.TrimSpace(ss[0])
			jobName = strings.TrimSpace(ss[1])
		}
		if jobId != "" {
			jobIds[jobId] = jobName
		}
	}
}
