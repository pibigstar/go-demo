package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func LoadYaml() {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	_, err := bundle.LoadMessageFile("zh-CN.yaml")
	if err != nil {
		panic(err)
	}
	localizer := i18n.NewLocalizer(bundle, "zh-CN")
	info := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "200",
		DefaultMessage: &i18n.Message{
			ID:    "200",
			Other: "未知错误",
		},
	})
	fmt.Println(info)
}

func LoadJson() {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFile("zh-CN.json")
	if err != nil {
		panic(err)
	}
	localizer := i18n.NewLocalizer(bundle, "zh-CN")
	info := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "200",
		DefaultMessage: &i18n.Message{
			ID:    "200",
			Other: "未知错误",
		},
	})
	fmt.Println(info)
}

func LoadJsonFile(bs []byte) {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustParseMessageFileBytes(bs, "zh.json")
	localizer := i18n.NewLocalizer(bundle, "zh-CN")
	info := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "200",
		DefaultMessage: &i18n.Message{
			ID:    "200",
			Other: "未知错误",
		},
	})
	fmt.Println(info)
}
