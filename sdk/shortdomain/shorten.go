package shortdomain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiURL = "https://api.uomg.com/api/long2dwz?%s"
)

type ShortResponse struct {
	Code     int    `json:"code"`
	ShortURL string `json:"ae_url"`
}

// https://www.free-api.com/doc/300
func GetShortURL(longURL string) (string, error) {

	params := url.Values{}
	params.Add("url", longURL)
	params.Add("dwzapi", "urlcn")

	resp, err := http.Get(fmt.Sprintf(apiURL, params.Encode()))
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	shortResponse := new(ShortResponse)
	err = json.Unmarshal(body, shortResponse)

	return shortResponse.ShortURL, err
}
