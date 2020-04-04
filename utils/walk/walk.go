package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// go build -ldflags="-H windowsgui"

func main() {
	var username, password *walk.TextEdit

	MainWindow{
		Title:  "QQ",
		Size:   Size{Width: 350, Height: 100},
		Layout: VBox{},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					HSplitter{
						Children: []Widget{
							TextLabel{Text: "用户名"},
							TextEdit{AssignTo: &username},
						},
					},
					HSplitter{
						Children: []Widget{
							TextLabel{Text: "密码"},
							TextEdit{AssignTo: &password},
						},
					},
				},
			},
			PushButton{
				Text: "登陆",
				OnClicked: func() {
					if username.Text() == "admin" && password.Text() == "admin" {
						walk.MsgBox(&walk.MainWindow{}, "成功", "登陆成功: "+username.Text(), walk.MsgBoxOK)
					}
				},
			},
		},
	}.Run()
}
