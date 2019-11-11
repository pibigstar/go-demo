package shortdomain

import "testing"

func TestGetShortURL(t *testing.T) {
	url, err := GetShortURL("https://pibigstar.github.io")
	if err != nil {
		t.Error(err)
	}
	t.Log(url)
}
