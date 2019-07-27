package oss

import (
	"fmt"
	"os"
	"testing"
)

func TestOSS(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = client.Put("oss/test.txt", file)
	if err != nil {
		fmt.Println(err.Error())
	}

	url := client.GetDownloadURL("oss/test.txt")
	fmt.Println(url)
}
