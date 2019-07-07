package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func qzone(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	url := r.Form.Get("data")
	saveInfo(url)
}

func read(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("qq")
	if err != nil {
		fmt.Println("read qq failed：", err.Error())
	}
	w.Write(data)
}

func main() {
	fmt.Println("启动成功。。。。。")
	http.HandleFunc("/qzone", qzone)         //设置访问的路由
	http.HandleFunc("/read", read)           //设置访问的路由
	err := http.ListenAndServe(":9500", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err.Error())
	}
}

func saveInfo(url string) {
	data := fmt.Sprintf("%s \n\n", url)
	file, err := os.OpenFile("qq", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create("qq")
	}
	defer file.Close()
	file.Write([]byte(data))
}
