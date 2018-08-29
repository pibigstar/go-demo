package get_post

import (
	"fmt"
	"go-demo/utils/errutil"
	"io/ioutil"
	"net/http"
)

func GetInfo() {

	response, err := http.Get("http://www.baidu.com")
	errutil.Check(err)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	errutil.Check(err)
	fmt.Println(string(data))
}
