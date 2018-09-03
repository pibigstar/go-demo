package dao

import (
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/xiocode/weigo"

	. "github.com/mnhkahn/maodou/models"
)

type WeiboConfig struct {
	Code      string
	Token     string
	ShortName string `json:"short_name"`
	Secret    string `json:"secret"`
	ThreadKey string `json:"thread_key"`
}

type WeiboDao struct {
}

// http://open.weibo.com/wiki/%E6%8E%88%E6%9D%83%E6%9C%BA%E5%88%B6%E8%AF%B4%E6%98%8E
// http://open.weibo.com/wiki/Help/error
func (this *WeiboDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	d := new(WeiboDaoContainer)
	config := new(WeiboConfig)
	err := json.Unmarshal([]byte(dsn), config)
	d.config = config
	if err != nil {
		return d, fmt.Errorf("Config for Weibo is error: %v", err)
	}

	d.Api = weigo.NewAPIClient("3925283339", "63dd249308e0638dab900781d0694131", "http://www.cyeam.com", "code")
	authorize_url, _ := d.Api.GetAuthorizeUrl(nil)
	fmt.Println("浏览器打开这个地址，输入code", authorize_url)
	// fmt.Scanf("%s", &d.config.Code)

	// var result map[string]interface{}
	// err = d.Api.RequestAccessToken(d.config.Code, &result) // code
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(result)
	// access_token := result["access_token"]
	// fmt.Println(access_token)
	// expires_in := result["expires_in"]
	// fmt.Println(int64(expires_in.(float64)))

	// d.Api.SetAccessToken(access_token.(string), int64(expires_in.(float64)))
	// d.Api.SetAccessToken("2.00g1SVIBbVFeRE30adcf7ec8MbsB6E", 157679999)
	d.Api.SetAccessToken("2.00TjCp_DbVFeREafad671208unnOqD", 639401)
	return d, nil
}

type WeiboDaoContainer struct {
	config   *WeiboConfig
	is_debug bool
	Api      *weigo.APIClient
}

func (this *WeiboDaoContainer) Debug(is_debug bool) {
}

// http://open.weibo.com/wiki/Statuses/upload_url_text#HTTP.E8.AF.B7.E6.B1.82.E6.96.B9.E5.BC.8F
func (this *WeiboDaoContainer) AddResult(p *Result) {
	log.Println(this.config.Token)
	log.Println(p)
	kws := map[string]interface{}{
		"status": fmt.Sprintf("【%s】%s %s", p.Title, p.Description, p.Link),
	}
	result := new(weigo.Status)
	err := this.Api.POST_statuses_update(kws, result)
	log.Println(err)
	log.Println(result)
}

func (this *WeiboDaoContainer) AddResults(p []Result) {

}

func (this *WeiboDaoContainer) DelResult(id interface{}) {

}

func (this *WeiboDaoContainer) DelResults(source string) {

}

func (this *WeiboDaoContainer) UpdateResult(p *Result) {

}

func (this *WeiboDaoContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *WeiboDaoContainer) GetResultById(id int) *Result {
	p := new(Result)
	return p
}

func (this *WeiboDaoContainer) GetResultByLink(url string) *Result {
	p := new(Result)
	return p
}

func (this *WeiboDaoContainer) GetResult(author, sort string, limit, start int) []Result {
	return nil
}

func (this *WeiboDaoContainer) IsResultUpdate(p *Result) bool {
	is_update := false
	return is_update
}

func (this *WeiboDaoContainer) Search(q string, limit, start int) (int, float64, []Result) {
	return 0, 0, nil
}

func init() {
	Register("weibo", &WeiboDao{})
}
