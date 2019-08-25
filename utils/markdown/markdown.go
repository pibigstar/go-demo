package markdown

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func Parse(input []byte) {
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	ioutil.WriteFile("index.html", html, 0666)
}
