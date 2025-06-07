//go:build go1.18
// +build go1.18

package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// 通过 go:embed 文件名，可以将该文件内容读入到变量bs中
//
//go:embed test.txt
var bs []byte

func TestEmbed(t *testing.T) {
	fmt.Println(string(bs))
}

// 通过embed挂载某个目录

//go:embed static/*
var staticFiles embed.FS

func TestEmbedDir(t *testing.T) {
	// 创建一个文件系统，指向嵌入的 static 目录
	staticFS := http.FS(staticFiles)
	fileServer := http.FileServer(staticFS)
	// 将所有请求路由到文件服务器
	http.Handle("/", http.StripPrefix("/", fileServer))
	if err := http.ListenAndServe("9090", nil); err != nil {
		panic(err)
	}
}
