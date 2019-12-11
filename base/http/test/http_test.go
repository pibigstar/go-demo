package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Body struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestHttpTest(t *testing.T) {
	// 路由分发器
	mux := &http.ServeMux{}
	// Get
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.FormValue("name"))
		w.Write([]byte(r.FormValue("name")))
	})
	// Post
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		var body Body
		if r.Body != nil {
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
		}
		t.Logf("%+v", body)
		body.Age = 20
		bs, _ := json.Marshal(body)
		w.Write(bs)
	})
	// Http Server, 模拟的http服务器
	ts := httptest.NewServer(mux)

	// Get test
	response, err := http.Get(ts.URL + "/get?name=pibigstar")
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	t.Log(string(bs))

	// Post test
	body := &Body{Name: "Hello"}
	bs, _ = json.Marshal(body)
	response, err = http.Post(ts.URL+"/post", "application/json", bytes.NewReader(bs))
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	bs, _ = ioutil.ReadAll(response.Body)
	t.Log(string(bs))
}
