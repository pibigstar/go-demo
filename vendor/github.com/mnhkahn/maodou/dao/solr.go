package dao

import (
	. "github.com/mnhkahn/maodou/models"
	// "encoding/json"
	"fmt"
	"github.com/mnhkahn/maodou/request/goreq"
	"net/url"
	"time"
)

type SolrDao struct {
	Dsn string
}

func (this *SolrDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	solr := new(SolrDaoContainer)
	solr.dsn = dsn
	solr.solr_req = goreq.Request{
		ContentType: "application/json",
		UserAgent:   "Cyeambot",
		Timeout:     time.Duration(5) * time.Second,
		// Compression: goreq.Gzip(),
	}
	solr.solr_req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	// goreq.SetConnectTimeout(time.Duration(5) * time.Second)
	return solr, nil
}

type SolrDaoContainer struct {
	dsn      string
	is_debug bool
	solr_req goreq.Request
}

func (this *SolrDaoContainer) Debug(is_debug bool) {
	this.solr_req.ShowDebug = is_debug
}

func (this *SolrDaoContainer) AddResult(p *Result) {
	// Skip duplicate items
	if this.GetResultByLink(p.Link) != nil {
		fmt.Println("duplicate", p.Link)
		return
	}

	this.solr_req.Method = "POST"
	this.solr_req.Uri = this.dsn + "/update"

	addSolr := new(AddSolr)
	addSolr.Add.CommitWithin = 1000
	addSolr.Add.Doc = *p
	// addSolr.Add.Overwrite = true

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *addSolr
	this.solr_req.QueryString = query

	_, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}
}

func (this *SolrDaoContainer) AddResults(p []Result) {

}

func (this *SolrDaoContainer) DelResult(id interface{}) {
	this.solr_req.Method = "Result"
	this.solr_req.Uri = this.dsn + "/update"

	delSolr := new(DelSolr)
	delSolr.Del.Query = fmt.Sprintf(`id:%v`, id)
	delSolr.Del.CommitWithin = 1000

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *delSolr
	this.solr_req.QueryString = query

	_, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}
}

func (this *SolrDaoContainer) DelResults(source string) {
	this.solr_req.Method = "Result"
	this.solr_req.Uri = this.dsn + "/update"

	delSolr := new(DelSolr)
	delSolr.Del.Query = fmt.Sprintf(`source:%s`, source)
	delSolr.Del.CommitWithin = 1000

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *delSolr
	this.solr_req.QueryString = query

	_, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}
}

func (this *SolrDaoContainer) UpdateResult(p *Result) {

}

func (this *SolrDaoContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *SolrDaoContainer) GetResultById(id int) *Result {
	p := new(Result)
	return p
}

func (this *SolrDaoContainer) GetResultByLink(u string) *Result {
	this.solr_req.Method = "GET"
	this.solr_req.Uri = this.dsn + "/select"

	query := url.Values{}
	query.Add("wt", "json")
	query.Add("q", fmt.Sprintf("link:%s", u))
	query.Add("start", fmt.Sprintf("%d", 0))
	query.Add("rows", fmt.Sprintf("%d", 1))
	this.solr_req.QueryString = query

	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	solr_Results := new(SolrResult)
	err = res.Body.FromJsonTo(solr_Results)
	if err != nil {
		panic(err)
	}
	if len(solr_Results.Response.Docs) > 0 {
		return &(solr_Results.Response.Docs[0])
	}
	return nil
}

func (this *SolrDaoContainer) GetResult(author, sort string, limit, start int) []Result {
	this.solr_req.Method = "GET"
	this.solr_req.Uri = this.dsn + "/select"

	query := url.Values{}
	query.Add("wt", "json")
	query.Add("q", fmt.Sprintf("author:%s", author))
	if sort != "" {
		query.Add("sort", sort)
	}
	query.Add("start", fmt.Sprintf("%d", start))
	query.Add("rows", fmt.Sprintf("%d", limit))
	this.solr_req.QueryString = query

	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	solr_Results := new(SolrResult)
	err = res.Body.FromJsonTo(solr_Results)
	if err != nil {
		panic(err)
	}
	return solr_Results.Response.Docs
}

func (this *SolrDaoContainer) IsResultUpdate(p *Result) bool {
	is_update := false
	return is_update
}

func (this *SolrDaoContainer) Search(q string, limit, start int) (int, float64, []Result) {
	this.solr_req.Method = "GET"
	this.solr_req.Uri = this.dsn + "/select"

	query := url.Values{}
	query.Add("wt", "json")
	query.Add("q", fmt.Sprintf("description:%s", q))
	query.Add("start", fmt.Sprintf("%d", start))
	query.Add("rows", fmt.Sprintf("%d", limit))
	query.Add("hl", "true")
	query.Add("hl.simple.pre", "<em>")
	query.Add("hl.simple.Result", "</em>")
	query.Add("hl.fl", "description")
	query.Add("hl.highlightMultiTerm", "true")
	// query.Add("sort", "figure desc, create_time desc")
	this.solr_req.QueryString = query

	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	solr_Results := new(SolrResult)
	err = res.Body.FromJsonTo(solr_Results)
	if err != nil {
		panic(err)
	}
	// for i := 0; i < len(solr_Results.Response.Docs); i++ {
	// 	solr_Results.Response.Docs[i].Description = solr_Results.Highlighting[solr_Results.Response.Docs[i].Link]["description"][0]
	// }
	return solr_Results.Response.NumFound, solr_Results.ResponseHeader.QTime, solr_Results.Response.Docs
}

func init() {
	Register("solr", &SolrDao{})
}

type AddSolr struct {
	Add struct {
		CommitWithin int    `json:"commitWithin"`
		Doc          Result `json:"doc"`
		Overwrite    bool   `json:"overwrite"`
	} `json:"add"`
}

type DelSolr struct {
	Del struct {
		Query        string `json:"query"`
		CommitWithin int    `json:"commitWithin"`
	} `json:"delete"`
}

type SolrResult struct {
	Response struct {
		Docs     []Result `json:"docs"`
		NumFound int      `json:"numFound"`
		Start    int      `json:"start"`
	} `json:"response"`
	ResponseHeader SolrResponseHeader             `json:"responseHeader"`
	Error          SolrError                      `json:"error"`
	Highlighting   map[string]map[string][]string `json:"highlighting"`
}

type SolrResponseHeader struct {
	QTime  float64 `json:"QTime"`
	Params struct {
		Indent string `json:"indent"`
		Q      string `json:"q"`
		Wt     string `json:"wt"`
	} `json:"params"`
	Status int `json:"status"`
}

type SolrError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
