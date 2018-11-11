package chain

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(content string)
	next(handler Handler,content string)
}

// 广告过滤
type AdHandler struct {
	handler Handler
}

func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content,"广告","**",1)
	fmt.Println(newContent)
	ad.next(ad.handler,newContent)
}

func (ad *AdHandler) next(handler Handler,content string) {
	if ad.handler!=nil {
		ad.handler.Handle(content)
	}
}
// 涉黄过滤 
type YellowHandler struct {
	handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content,"涉黄","**",1)
	fmt.Println(newContent)
	yellow.next(yellow.handler,newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if yellow.handler!=nil {
		yellow.handler.Handle(content)
	}
}

// 敏感词过滤
type SensitiveHandler struct {
	handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤。。。")
	newContent := strings.Replace(content,"敏感词","***",1)
	fmt.Println(newContent)
	sensitive.next(sensitive.handler,newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if sensitive.handler!=nil {
		sensitive.handler.Handle(content)
	}
}








