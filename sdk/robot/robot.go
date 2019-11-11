package robot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	_RobotURL = "https://oapi.dingtalk.com/robot/send?access_token=a3c088934af0b084d1d7f18f98ffacfcdd18f560752711e2d37044b943bc21d4"
)

type Content struct {
	Level  string `json:"level,omitempty"`
	Method string `json:"method,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Svc    string `json:"svc,omitempty"`
	Code   int    `json:"code,omitempty"`
	Tid    string `json:"tid,omitempty"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Btns           []Btn  `json:"btns"`
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type atMobiles struct {
	phones []string `json:"phones"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicURL     string `json:"PicUrl"`
	MessageURL string `json:"messageUrl"`
}

type RobotMsg struct {
	Msgtype    string     `json:"msgtype"`
	Markdown   Markdown   `json:"markdown"`
	ActionCard ActionCard `json:"actionCard"`
	Link       Link       `json:"link"`
	At         atMobiles  `json:"at"`
	IsAtAll    bool       `json:"isAtAll"`
}

func Send(msg *RobotMsg) error {

	data, _ := json.Marshal(msg)

	resp, err := http.Post(_RobotURL, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
