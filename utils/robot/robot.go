package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

// 需要安装 MinGW
// http://store.pibigstar.com/mingw-w64-install.exe
func main() {
	go ListenScreen()

	go ListenKeywords()

	ShowMessage()

	time.Sleep(10 * time.Second)

	// 模拟按键
	PressKeywords()
	// 杀死进程
	FindAndKillProcess("robot")
}

// 监听屏幕
func ListenScreen() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
			robotgo.FindBitmap(bitmap)
			robotgo.SaveBitmap(bitmap, "test.png")
		}
	}
}

// 监听键盘事件
func ListenKeywords() {
	keywords := []string{"A", "B", "C", "D"}
	for _, key := range keywords {
		go func(key string) {
			ok := robotgo.AddEvent(key)
			if ok {
				fmt.Printf("press: %s \n", key)
			}
		}(key)
	}
}

// 模拟键盘事件
func PressKeywords() {
	// H
	robotgo.KeyTap("H")
	// 回车
	robotgo.KeyTap("enter")
	// 键入文字
	robotgo.TypeStr("Hello World")
	// 关掉程序
	robotgo.KeyTap("c", "ctrl")
}

// 弹窗
func ShowMessage() {
	btMsg := robotgo.ShowAlert("Alarm", "Hello Pibigstar", "Success", "Close")
	if btMsg == 0 {
		fmt.Println("确定")
	} else {
		fmt.Println("取消")
	}
}

// 找到并杀死进程
func FindAndKillProcess(name string) {
	ids, err := robotgo.FindIds(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, id := range ids {
		err := robotgo.Kill(id)
		if err == nil {
			fmt.Printf("杀死进程: %d", id)
		}
	}
}
