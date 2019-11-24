package robot

import (
	"fmt"
	"go-demo/utils/env"
	"strings"
	"testing"
	"time"
)

const (
	elkDev = "http://devops.izhaohu.com/kibana/app/kibana#/discover?_g=(refreshInterval:(pause:!t,value:0),time:(from:now/w,to:now/w))&_a=(columns:!(_source),index:'5d5cba20-db77-11e9-99de-0fa45564f0e7',interval:auto,query:(language:kuery,query:'%s'),sort:!(!('@timestamp',desc)))"
)

func TestActionCard(t *testing.T) {
	if env.IsCI() {
		return
	}

	btn := []Btn{
		{
			Title:     "查看详细内容",
			ActionURL: fmt.Sprintf(elkDev, "623b64ad171c248363e491d46d65b7ca"),
		},
	}

	actionCard := ActionCard{
		Title:          "发生特殊异常",
		Text:           buildMsg(),
		HideAvatar:     "0",
		BtnOrientation: "0",
		Btns:           btn,
	}

	msg := &RobotMsg{
		Msgtype:    "actionCard",
		ActionCard: actionCard,
		IsAtAll:    false,
	}

	Send(msg)
}

func TestSendMarkDown(t *testing.T) {
	if env.IsCI() {
		return
	}

	markdown := Markdown{
		Title: "发生特殊异常",
		Text:  buildMsg(),
	}

	msg := &RobotMsg{
		Msgtype:  "markdown",
		Markdown: markdown,
		IsAtAll:  false,
	}

	Send(msg)
}

func TestSendLink(t *testing.T) {
	if env.IsCI() {
		return
	}
	link := Link{
		Title:      "发生特殊",
		Text:       buildMsg(),
		PicURL:     "",
		MessageURL: fmt.Sprintf(elkDev, "623b64ad171c248363e491d46d65b7ca"),
	}
	msg := &RobotMsg{
		Msgtype: "link",
		Link:    link,
	}
	Send(msg)
}

func buildMsg() string {
	contents := []*Content{
		&Content{
			Level:  "Dev",
			Method: "biz.SalaryService/RunJob",
			Msg:    "特殊异常",
			Svc:    "biz",
			Tid:    "123456789",
		},
		&Content{
			Level:  "Dev",
			Method: "biz.AttendanceService/RunJob",
			Msg:    "特殊异常",
			Svc:    "biz",
			Tid:    "123456789",
		},
	}

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("#### 环境【<font color=#FF0000>%s</font>】%s \n\n", "dev", time.Now().Format("2006-01-02 15:04:05")))
	for _, content := range contents {
		str.WriteString("------------------- \n\n")
		str.WriteString(fmt.Sprintf("##### Svc: %s \n", content.Svc))
		str.WriteString(fmt.Sprintf("##### Method: %s \n", content.Method))
		str.WriteString(fmt.Sprintf("##### Code: %d \n", content.Code))
		str.WriteString(fmt.Sprintf("##### Tid: %s \n", content.Tid))
		str.WriteString(fmt.Sprintf("##### Msg: %s \n", content.Msg))
	}
	return str.String()
}
