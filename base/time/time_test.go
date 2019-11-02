package time

import (
	"fmt"
	"testing"
	"time"
)

const defaultTime = "2006-01-02 15:04:05"

// 时间格式化
func TimeFormat(date time.Time) string {
	nowTime := date.Format(defaultTime)
	return nowTime
}

func TestTimeParse(t *testing.T) {
	now := time.Now()
	// 时间格式化
	format := TimeFormat(now)
	t.Log(format)
	// 字符串转时间
	parse, err := time.Parse(defaultTime, format)
	if err != nil {
		t.Error(err)
	}
	// 三小时后
	parse.Add(time.Hour * 3)
	// 一小时前
	parse.Add(time.Hour * -1)
	// 获取时间戳
	t1 := parse.Unix()
	// 获取纳秒级时间戳
	t2 := parse.UnixNano()
	t.Log(t1, t2)
}

func TestTimeout(t *testing.T) {
	done := make(chan struct{})
	deadLine := time.Now().Add(time.Second * 1)

	dur := time.Until(deadLine)
	if dur <= 0 {
		fmt.Println("当前已超过deadLine,停止方法")
	}

	time.AfterFunc(dur, func() {
		fmt.Println("到达deadLine，停止方法")
		done <- struct{}{}

	})
	select {
	case <-done:
	}
}
