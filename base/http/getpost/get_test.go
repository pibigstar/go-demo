package getpost

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
