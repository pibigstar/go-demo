package markdown

import (
	"github.com/TruthHun/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func Parse(input []byte) {
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	ioutil.WriteFile("index.html", html, 0666)
}

func htmlToMarkdown(html string) string {
	mdStr := html2md.Convert(html)
	return mdStr
}
