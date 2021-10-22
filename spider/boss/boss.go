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
	"net/smtp"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	communicatedNo  = 1 // 未被同事沟通过
	communicatedYes = 2 // 已被同事沟通过

	requestTypeToMe   = "3" // 牛人向我发简历
	requestTypeToGeek = "4" // 我向牛人发求简历
)

var (
	jobIds  = make(map[string]string) // 工作Id
	talked  sync.Map                  // 已经沟通过的人
	logFile *os.File

	client = http.Client{
		Timeout: 5 * time.Second,
	}

	maxLimit       = errors.New("今日沟通已达上限")
	notFriend      = errors.New("好友关系校验失败")
	notLogin       = errors.New("当前登录状态已失效")
	accountUnusual = errors.New("账号异常")

	school985   []string
	school211   []string
	goodCompany []string
	cookie      string

	cookieFile    = "cookie.txt"
	school985File = "985.txt"
	school211File = "211.txt"
	jobsFile      = "jobs.txt"
	companyFile   = "company.txt"
	bossLog       = "boss.log"

	//log = logs.New(logFile, "", logs.Ldate|logs.Ltime)
	log = logs.New(os.Stdout, "", logs.Ldate|logs.Ltime)
)

func initBoss() {
	// 设置当前运行目录
	setFilePath()
	// 读取cookie信息
	readCookie()
	// 监听cookie文件
	watchCookie()
	// 读取jobId
	readJobs()
	// 读取学校信息
	readSchool()
	// 读取大厂信息
	readCompany()
	// 设置自动打招呼语
	setHelloMsg()
}

func main() {

	initBoss()

	if len(jobIds) == 0 {
		inputJobs()
	}
	if len(jobIds) == 0 {
		fmt.Println("暂时没有需要沟通的职位~")
		return
	}

	var wg sync.WaitGroup
	for jobId, jobName := range jobIds {
		wg.Add(1)
		fmt.Println("正在沟通职位:", jobName)
		go func(jobId, jobName string) {
			defer wg.Done()
			defer func() {
				if e := recover(); e != nil {
					log.Println("recover", e)
				}
				Hiring(jobId, jobName)
			}()
		}(jobId, jobName)
	}
	wg.Wait()
}

// 输入存储job信息
func inputJobs() {
	if len(jobIds) > 0 {
		return
	}
	jobs := listJobs()
	if len(jobs) == 0 {
		fmt.Println("你没有开放的职位")
		return
	}
	for i, j := range jobs {
		fmt.Printf("编号:%d 职位: %s \n", i, j.JobName)
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入你要沟通的职位编号:")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("输入错误!")
		return
	}
	ids := strings.Split(input, ",")
	var str string
	for _, id := range ids {
		id = strings.ReplaceAll(id, "\n", "")
		id = strings.TrimSpace(id)
		i, err := strconv.Atoi(id)
		if err != nil || i >= len(jobs) {
			fmt.Printf("%s 编号有误 \n", id)
			continue
		}
		jobId := jobs[i].JobId
		jobName := jobs[i].JobName
		jobIds[jobId] = jobName
		str += fmt.Sprintf("%s   //%s \n", jobId, jobName)
	}

	// 储存
	err = ioutil.WriteFile(jobsFile, []byte(str), 0666)
	if err != nil {
		log.Println("存储jobId信息失败", err.Error())
	}
}

// 招人
func Hiring(jobId, jobName string) {
	var (
		// 10秒一次，防止被反爬
		t = time.NewTicker(10 * time.Second)
		// 10分钟候选人选择
		ctx, cancel = context.WithTimeout(context.Background(), time.Minute*3)
		geeksQueue  []*Geek
	)
	defer cancel()

	for {
		select {
		case <-t.C:
			geeks, err := searchGeekByJobId(jobId, jobName)
			if err != nil {
				if err == notLogin {
					sendFeiShu("Boss当前登录状态失效")
				}
				// 通知可以去打招呼了
				cancel()
				t.Stop()
			}
			geeksQueue = append(geeksQueue, geeks...)
			// 有10个候选人就没必要继续跑了
			if len(geeksQueue) == 10 {
				cancel()
				t.Stop()
			}

		case <-ctx.Done():
			// 打招呼并请求简历
			helloAndRequestResumes(jobId, geeksQueue)
			return
		}
	}
}

// 打招呼并轮询请求简历
func helloAndRequestResumes(jobId string, geeksQueue []*Geek) {
	// 按权重排序
	sort.Sort(SortGeek(geeksQueue))
	var wg sync.WaitGroup
	// 进行一小时的请求简历
	rrCtx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	for _, l := range geeksQueue {
		if _, ok := talked.Load(l.GeekCard.GeekID); ok {
			continue
		}
		err := hello(jobId, l.GeekCard.EncryptGeekID, l.GeekCard.Lid, l.GeekCard.SecurityID, l.GeekCard.ExpectID)
		if err != nil {
			if err == maxLimit {
				log.Println("今日已达上限")
				break
			}
			log.Printf("打招呼失败，候选人: %s, 分值: %d err: %s\n", l.GeekCard.GeekName, l.Weight, err.Error())
			continue
		}

		log.Printf("正在与: %s 打招呼, 分值: %d\n", l.GeekCard.GeekName, l.Weight)

		// 标记已经打过招呼了
		talked.Store(l.GeekCard.GeekID, 1)

		// 轮询向牛人直接请求简历直到对方回复我们建立好友关系
		wg.Add(1)
		go func(name, securityId string) {
			defer wg.Done()
			t := time.NewTicker(time.Minute * 5)
			for {
				select {
				case <-t.C:
					log.Printf("正在索求候选人:%s的简历 \n", name)
					if err := requestResumes(name, securityId); err == nil {
						t.Stop()
						return
					}
				case <-rrCtx.Done():
					t.Stop()
					return
				}
			}
		}(l.GeekCard.GeekName, l.GeekCard.SecurityID)

		time.Sleep(5 * time.Second)
	}

	wg.Wait()

	log.Println("====本次招聘任务结束====")
}

func searchGeekByJobId(jobId, jobName string) ([]*Geek, error) {
	var geeks []*Geek
	geekList, err := listRecommend(jobId)
	if err != nil {
		return nil, err
	}
	if len(geekList) == 0 {
		sendFeiShu("Boss当前需要重新验证")
		return nil, accountUnusual
	}

	for _, geek := range geekList {
		log.Printf("候选人: %s  期待职位：%s \n", geek.GeekCard.GeekName, geek.GeekCard.ExpectPositionName)
		if selectGeek(geek, jobName) {
			// 分高的直接去打招呼，防止中途被反爬导致程序提前退出
			if geek.Weight >= 15 {
				go func(geek *Geek) {
					err := hello(jobId, geek.GeekCard.EncryptGeekID, geek.GeekCard.Lid, geek.GeekCard.SecurityID, geek.GeekCard.ExpectID)
					if err != nil {
						log.Printf("候选人: %s打招呼失败, 分值: %d\n", geek.GeekCard.GeekName, geek.Weight)
					}
				}(geek)
				continue
			}
			// 分值低于 15 的才需要进行比较
			log.Printf("候选人: %s  进入队列, 分值: %d\n", geek.GeekCard.GeekName, geek.Weight)
			geeks = append(geeks, geek)
		}
	}
	return geeks, nil
}

// 筛选并打分
func selectGeek(geek *Geek, jobName string) bool {
	// 已经打过招呼了
	if geek.HaveChatted == 1 {
		return false
	}
	// 已经被同事撩过
	if geek.Cooperate == communicatedYes {
		return false
	}
	// 岗位匹配
	if !matchJob(jobName, geek) {
		return false
	}
	//  是否是本科
	if geek.GeekCard.GeekDegree == "本科" {
		geek.Weight += 3
	}
	//  是否是硕士
	if geek.GeekCard.GeekDegree == "硕士" {
		geek.Weight += 4
	}
	// 是否是211
	if isContains(school211, geek.GeekCard.GeekEdu.School) {
		geek.Weight += 3
	}
	// 是否是985
	if isContains(school985, geek.GeekCard.GeekEdu.School) {
		geek.Weight += 4
	}
	// 是否在大厂
	for _, w := range geek.GeekCard.GeekWorks {
		if isContains(goodCompany, w.Company) {
			geek.Weight += 5
			break
		}
	}
	// 工作年限大于3年
	workStr := strings.ReplaceAll(geek.GeekCard.GeekWorkYear, "年", "")
	if years, err := strconv.Atoi(workStr); err == nil && years >= 3 {
		geek.Weight += 3
	}
	// 年龄
	ageStr := strings.ReplaceAll(geek.GeekCard.AgeDesc, "岁", "")
	if age, err := strconv.Atoi(ageStr); err == nil && age >= 26 && age <= 35 {
		geek.Weight += 2
	}
	// 在职-月内到岗
	if strings.Contains(geek.GeekCard.ApplyStatusDesc, "月内到岗") {
		geek.Weight += 2
	}
	// 离职-随时到岗
	if strings.Contains(geek.GeekCard.ApplyStatusDesc, "离职") {
		geek.Weight += 3
	}
	// 今日活跃
	if strings.Contains(geek.ActiveTimeDesc, "今日活跃") {
		geek.Weight += 1
	}
	// 刚刚活跃
	if strings.Contains(geek.ActiveTimeDesc, "刚刚活跃") {
		geek.Weight += 2
	}
	return true
}

// 岗位是否匹配
func matchJob(jobName string, geek *Geek) bool {
	jobName = strings.ToLower(jobName)
	expectPositionName := strings.ToLower(geek.GeekCard.ExpectPositionName)
	// 期望职位是否匹配（todo: 这个不准）
	if strings.Contains(jobName, expectPositionName) || strings.Contains(expectPositionName, jobName) {
		return true
	}
	// 个人描述里面是否有该岗位
	if strings.Contains(geek.GeekCard.GeekDesc.Content, jobName) {
		return true
	}
	// 历史工作里面是否有该岗位
	for _, j := range geek.GeekCard.GeekWorks {
		j.PositionName = strings.ToLower(j.PositionName)
		if strings.Contains(j.PositionName, jobName) || strings.Contains(jobName, j.PositionName) {
			return true
		}
	}
	return false
}

func isContains(arrs []string, arr string) bool {
	for _, s := range arrs {
		if strings.EqualFold(s, arr) {
			return true
		}
		if strings.Contains(arr, s) {
			return true
		}
		if strings.Contains(s, arr) {
			return true
		}
	}
	return false
}

// 打招呼
// 需要设置自动打招呼
func hello(jobId, encryptGeekId, lid, securityId string, expectId int) error {
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
func requestResumes(name, securityId string) error {
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
	if strings.Contains(string(bs), "好友关系校验失败") {
		return notFriend
	}
	var temp *RequestResumesResp
	if err = json.Unmarshal(bs, &temp); err != nil {
		return err
	}
	if fmt.Sprintf("%d", temp.ZpData.Type) != requestTypeToGeek {
		return notFriend
	}
	log.Printf("请求候选人:%s的简历成功 \n", name)
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

	resp, err := client.Do(req)
	if err != nil {
		log.Println("ListRecommend request", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bs), "登录状态已失效") {
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
	req.Header.Add("cookie", cookie)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("upgrade-insecure-requests", "1")
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

func readCompany() {
	bs, _ := ioutil.ReadFile(companyFile)
	br := bufio.NewReader(bytes.NewReader(bs))
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		goodCompany = append(goodCompany, string(a))
	}
}

func readJobs() {
	bs, err := ioutil.ReadFile(jobsFile)
	if err != nil {
		return
	}
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

// 设置自动打招呼语
// 根据Job设置
func setHelloMsg() {
	// 开启自动打招呼
	uri := "https://www.zhipin.com/wapi/zpchat/greeting/updateGreeting"
	values := url.Values{}
	values.Add("status", "1")
	values.Add("templateId", "")
	req, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
	addHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("open auto greeting", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bs), "Success") {
		log.Println("已开启自动打招呼")
	}
	// 获取职位列表
	uri = "https://www.zhipin.com/wapi/zpchat/greeting/job/get"
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	addHeader(req)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println("setHelloMsg get", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ = ioutil.ReadAll(resp.Body)
	var t *JobHelloMsg
	err = json.Unmarshal(bs, &t)
	if err != nil {
		log.Println("unmarshal job message", err.Error())
		return
	}

	// 设置每个岗位的打招呼语
	uri = "https://www.zhipin.com/wapi/zpchat/greeting/job/save"
	for _, job := range t.ZpData.Jobs {
		// 如果设置过了,就不再设置了
		if job.JobGreeting != "" {
			continue
		}
		data := url.Values{}
		data.Add("encJobId", job.EncJobID)
		data.Add("encGreetingId", job.EncGreetingID)
		data.Add("content", fmt.Sprintf("你好，这边是得物APP，我们目前正在大力扩招%s，如果您有兴趣的话，方便发一份简历给我吗？期待你的加入～", job.JobName))

		req, _ = http.NewRequest(http.MethodPost, uri, strings.NewReader(data.Encode()))
		addHeader(req)
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			log.Println("save job hell msg", err.Error())
			continue
		}
		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(bs), "Success") {
			log.Printf("设置职位: %s 的打招呼语成功", job.JobName)
		}
	}
}

func sendEmail() {
	var (
		username = "741047261@qq.com"
		password = ""
		host     = "smtp.qq.com"
		addr     = "smtp.qq.com:25"
	)
	auth := smtp.PlainAuth("", username, password, host)

	user := "741047261@qq.com"
	to := []string{"741047261@qq.com"}
	msg := []byte(`From: 741047261@qq.com
To: 741047261@qq.com
Subject: Boss登录状态失效

boss登录状态已失效，请及时更改
`)
	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		log.Println("发送邮件提醒失败:", err.Error())
	}
}

//  扫码登录
func getQRId(ctx context.Context) {
	// 取qrId
	uri := "https://login.zhipin.com/wapi/zppassport/captcha/randkey"
	values := url.Values{}
	values.Add("pk", "cpc_user_sign_up")
	req, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
	addHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("get qr id", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	var msg *QRMsg
	if err = json.Unmarshal(bs, &msg); err != nil {
		log.Println("unmarshal qr msg", err.Error())
		return
	}
	// 取qrId
	qrId := msg.ZpData.QrID

	newCtx, _ := context.WithTimeout(ctx, 10*time.Minute)
	go func(qrId string) {
		t := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-t.C:
				// 获取set-cookie
				if err := setCookie(qrId); err == nil {
					t.Stop()
					return
				}
			case <-newCtx.Done():
				return
			}
		}

	}(qrId)

}

func setCookie(qrId string) error {
	uri := "https://login.zhipin.com/wapi/zppassport/qrcode/dispatcher?"
	values := url.Values{}
	values.Add("qrId", qrId)
	values.Add("_", fmt.Sprintf("%d", time.Now().Unix()))
	req, _ := http.NewRequest(http.MethodGet, uri, strings.NewReader(values.Encode()))
	addHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("get cookie", err.Error())
		return err
	}
	defer resp.Body.Close()

	var setCookie string
	for _, c := range resp.Header["Set-Cookie"] {
		setCookie += c
		setCookie += ";"
	}
	if setCookie == "" {
		return fmt.Errorf("no cookie")
	}
	return nil
}

// 获取job列表
func listJobs() []*Job {
	uri := "https://www.zhipin.com/wapi/zpjob/job/data/list?"
	values := url.Values{}
	values.Add("position", "0")
	values.Add("searchStr", "0")
	values.Add("page", "1")
	values.Add("_", fmt.Sprintf("%d", time.Now().Unix()))

	req, _ := http.NewRequest(http.MethodGet, uri, strings.NewReader(values.Encode()))
	addHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("list job", err.Error())
		return nil
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	var jResp *JobListResp
	if err = json.Unmarshal(bs, &jResp); err != nil {
		log.Println("unmarshal list job", err.Error())
		return nil
	}
	var jobs []*Job
	for _, j := range jResp.ZpData.Data {
		if j.JobStatus == 0 {
			jobs = append(jobs, &Job{
				JobId:   j.EncryptJobID,
				JobName: j.JobName,
			})
		}
	}
	return jobs
}

func setFilePath() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(currentFile)

	cookieFile = filepath.Join(basePath, cookieFile)
	school985File = filepath.Join(basePath, school985File)
	school211File = filepath.Join(basePath, school211File)
	jobsFile = filepath.Join(basePath, jobsFile)
	companyFile = filepath.Join(basePath, companyFile)
	logFile, _ = os.OpenFile(filepath.Join(basePath, bossLog), os.O_RDWR|os.O_CREATE, 0664)
}

func sendFeiShu(msg string) {
	uri := "https://open.feishu.cn/open-apis/bot/v2/hook/9b46f934-2e77-499e-81e8-af02b4b27cde"
	text := fmt.Sprintf(`
	{
		"msg_type": "text",
		"content": {
			"text": "%s"
		}
	}`, msg)

	_, err := http.Post(uri, "application/json", strings.NewReader(text))
	if err != nil {
		log.Println("发送飞书提醒失败, msg:", msg)
	}
}
