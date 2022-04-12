//go:build go1.18
// +build go1.18

package embed

import (
	_ "embed"
	"fmt"
	"testing"
)

// 通过 go:embed 文件名，可以将该文件内容读入到变量bs中
//go:embed test.txt
var bs []byte

func TestEmbed(t *testing.T) {
	fmt.Println(string(bs))
}
