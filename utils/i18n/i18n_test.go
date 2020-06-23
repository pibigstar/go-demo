package i18n

import (
	"go-demo/utils/i18n/locales"
	"io/ioutil"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	LoadYaml()
}

func TestLoadJson(t *testing.T) {
	LoadJson()
}

func TestLoadJsonFile(t *testing.T) {
	filePath := locales.Path("zh-CN.json")
	f, _ := ioutil.ReadFile(filePath)
	LoadJsonFile(f)
}
