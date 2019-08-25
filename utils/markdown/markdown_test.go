package markdown

import (
	"io/ioutil"
	"testing"
)

func TestMarkdown(t *testing.T) {
	bytes, err := ioutil.ReadFile("index.md")
	if err != nil {
		t.Error(err)
	}
	Parse(bytes)
}
