package shortdomain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiShortURL = "https://dwz.cn/admin/v2/create"
	shortToken  = "fa44d0ecaabdf1a3629faad123f1a50e"
)

type ShortResponse struct {
	Code     int    `json:"Code"`
	ShortUrl string `json:"ShortUrl"`
	LongUrl  string `json:"LongUrl"`
	ErrMsg   string `json:"ErrMsg"`
}

func GetShortURL(longURL string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiShortURL, strings.NewReader("url="+longURL))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Token", shortToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	shortResponse := new(ShortResponse)
	json.Unmarshal(body, shortResponse)
	fmt.Printf("%+v", shortResponse)

	return shortResponse.ShortUrl, nil
}
