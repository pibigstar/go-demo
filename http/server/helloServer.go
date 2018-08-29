package server

import (
	"fmt"
	"net/http"
)

// 创建一个简单的web服务器
func HelloServer() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8088", nil)
}

func HelloHandler(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	name := request.Form.Get("name")

	w.Write([]byte(fmt.Sprintf("<h1>hello %s</h1>", name)))
}

// 创建一个简单的文件服务器
func FileServer() {
	http.ListenAndServe(":8089", http.FileServer(http.Dir(".")))
}
