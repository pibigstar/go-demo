package sdk

import (
	"io/ioutil"
	"testing"
)

func TestGenQRCode(t *testing.T) {
	bytes, err := GenDefaultQRCode("Hello World")
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("1.jpg", bytes, 0666)

	bytes, err = GenDefaultQRCode("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("2.jpg", bytes, 0666)
}
