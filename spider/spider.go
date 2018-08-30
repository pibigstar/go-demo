package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/cygo"
	"github.com/mnhkahn/maodou/dao"
	"github.com/mnhkahn/maodou/models"
	"log"
	"strconv"
	"strings"
)

type Haixiu struct {
	maodou.MaoDou
}

func (this *Haixiu) Start() {
	response, err := this.Cawl("http://www.douban.com/group/haixiuzu/discussion")
	if err != nil {
		log.Println("Error when cawl", err.Error())
	}
	this.Index(response)
}

func (this *Haixiu) Index(resp *maodou.Response) {
	resp.Doc(`#content > div > div.article > div:nth-child(2) > table > tbody > tr > td.title > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			res, err := this.Cawl(href)
			if err != nil {
				log.Println("Error when cawl", err)
			}
			this.Detail(res)
		}
	})
}

func (this *Haixiu) Detail(resp *maodou.Response) {
	res := new(models.Result)
	res.Id = strconv.ParseUint(strings.Split(resp.Url, "/")[5], 64, 11)
	res.Title = resp.Doc("#content > h1").Text()
	res.Author = resp.Doc("#content > div > div.article > div.topic-content.clearfix > div.topic-doc > h3 > span.from > a").Text()
	res.Figure, _ = resp.Doc("#link-report > div.topic-content > div.topic-figure.cc > img").Attr("src")
	res.Link = resp.Url
	res.Source = "www.douban.com/group/haixiuzu/discussion"
	res.ParseDate = cygo.Now()
	res.Description = "haixiuzu"
	this.Result(res)
}

func (this *Haixiu) Result(result *models.Result) {
	if result.Figure != "" {
		Dao, err := dao.NewDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d"}`)
		if err != nil {
			panic(err)
		}
		Dao.AddResult(result)
	}
}

func main() {
	maodou.Register(new(Haixiu), 30)
}
