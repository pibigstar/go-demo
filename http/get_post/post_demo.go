package get_post

import (
	"fmt"
	"go-demo/utils/errutil"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostData() {

	response, err := http.PostForm("http://localhost:8080/partner/login", url.Values{"username": {"rob"}, "password": {"abc123_"}})
	errutil.Check(err)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	errutil.Check(err)

	fmt.Println(string(data))

}
