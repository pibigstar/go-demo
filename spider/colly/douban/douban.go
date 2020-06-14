package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

const URL = "https://movie.douban.com/top250"

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.IgnoreRobotsTxt(),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"),
	)

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		title := e.DOM.Find("span.title").Eq(0).Text()
		url, _ := e.DOM.Find("a").Eq(0).Attr("href")
		fmt.Println(title, url)
	})

	c.OnHTML(".next", func(e *colly.HTMLElement) {
		next, b := e.DOM.Find("a").Attr("href")
		if b {
			nextURL := fmt.Sprintf("%s%s", URL, next)
			e.Request.Visit(nextURL)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(res *colly.Response) {
		// 处理Response
	})

	c.OnError(func(res *colly.Response, err error) {
		// 处理Error
		fmt.Println("err", err.Error())
	})

	c.Visit(URL)
	c.Wait()
}
