package dao

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mnhkahn/maodou/request/goreq"

	. "github.com/mnhkahn/maodou/models"
)

type DuoShuoConfig struct {
	ShortName string `json:"short_name"`
	Secret    string `json:"secret"`
	ThreadKey string `json:"thread_key"`
}

type DuoShuoDao struct {
}

func (this *DuoShuoDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	d := new(DuoShuoDaoContainer)
	config := new(DuoShuoConfig)
	err := json.Unmarshal([]byte(dsn), config)
	d.config = config
	if err != nil {
		return d, fmt.Errorf("Config for duoshuo is error: %v", err)
	}

	goreq.SetConnectTimeout(5 * time.Second)

	return d, nil
}

type DuoShuoDaoContainer struct {
	config   *DuoShuoConfig
	is_debug bool
	req      goreq.Request
}

func (this *DuoShuoDaoContainer) Debug(is_debug bool) {
	// this.req.ShowDebugDetail = is_debug
}

func (this *DuoShuoDaoContainer) AddResult(p *Result) {
	this.req.Method = "POST"
	this.req.Uri = "http://api.duoshuo.com/posts/import.json"
	this.req.ContentType = "application/x-www-form-urlencoded"
	this.req.Timeout = time.Duration(60) * time.Second

	duoshuo_byte, _ := json.Marshal(*p)
	this.req.Body = fmt.Sprintf("short_name=%s&secret=%s&posts[0][post_key]=%s&posts[0][thread_key]=%s&posts[0][message]=%s", this.config.ShortName, this.config.Secret, p.Id, this.config.ThreadKey, base64.URLEncoding.EncodeToString(duoshuo_byte))
	resp, err := this.req.Do()
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		var err_str string
		if resp != nil {
			err_str, _ = resp.Body.ToString()
			err_str = fmt.Sprintf("%d %s", resp.StatusCode, err_str)
		}
		log.Printf("Error: %s\n", err_str)
	} else {
		log.Println("Add to DuoShuo Success.")
	}
}

func (this *DuoShuoDaoContainer) AddResults(p []Result) {

}

func (this *DuoShuoDaoContainer) DelResult(id interface{}) {

}

func (this *DuoShuoDaoContainer) DelResults(source string) {

}

func (this *DuoShuoDaoContainer) UpdateResult(p *Result) {

}

func (this *DuoShuoDaoContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *DuoShuoDaoContainer) GetResultById(id int) *Result {
	p := new(Result)
	return p
}

func (this *DuoShuoDaoContainer) GetResultByLink(url string) *Result {
	p := new(Result)
	return p
}

func (this *DuoShuoDaoContainer) GetResult(author, sort string, limit, start int) []Result {
	return nil
}

func (this *DuoShuoDaoContainer) IsResultUpdate(p *Result) bool {
	is_update := false
	return is_update
}

func (this *DuoShuoDaoContainer) Search(q string, limit, start int) (int, float64, []Result) {
	// this.req.Method = "GET"
	// this.req.Uri = "http://api.duoshuo.com/threads/listResults.json"
	// this.req.ContentType = "application/x-www-form-urlencoded"

	// addDuoShuo := url.Values{}
	// addDuoShuo.Add("short_name", this.config.ShortName)
	// addDuoShuo.Add("secret", this.config.Secret)
	// addDuoShuo.Add("Results[0][Result_key]", p.Id)
	// addDuoShuo.Add("Results[0][thread_key]", "haixiuzu-cyeam")

	// duoshuo_byte, _ := json.Marshal(addDuoShuo)
	// addDuoShuo.Add("Results[0][message]", base64.URLEncoding.EncodeToString(duoshuo_byte))
	// this.req.Body = addDuoShuo.Encode()

	// resp, err := this.req.Do()
	// if err != nil {
	// 	panic(err)
	// }
	// if resp.StatusCode != 200 {
	// 	err_str, _ := resp.Body.ToString()
	// 	panic(err_str)
	// }
	return 0, 0, nil
}

func init() {
	Register("duoshuo", &DuoShuoDao{})
}
