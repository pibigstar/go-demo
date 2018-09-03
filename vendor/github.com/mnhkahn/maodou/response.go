package maodou

import (
	"errors"
	"golang.org/x/net/html"
	"io"

	"github.com/PuerkitoBio/goquery"

	"github.com/mnhkahn/maodou/parser"
)

type Response struct {
	Url         string
	RawDocument *html.Node
	Document    *goquery.Document
}

func NewResponse(r io.Reader, url string) (*Response, error) {
	var err error
	if r == nil {
		return nil, errors.New("Error Reader.")
	}
	resp := new(Response)
	resp.Url = url
	resp.RawDocument, err = html.Parse(r)
	resp.Document = goquery.NewDocumentFromNode(resp.RawDocument)
	return resp, err
}

func (this *Response) Content(is_optimizatioin bool) string {
	return parser.ContentFromNode(this.RawDocument, is_optimizatioin)
}

func (this *Response) Doc(css string) *goquery.Selection {
	return this.Document.Find(css)
}
