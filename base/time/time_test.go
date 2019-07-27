package time

import (
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
		t.Fatal(err)
	}
	// 三小时后
	parse.Add(time.Hour * 3)
	// 一小时前
	parse.Add(time.Hour * -1)
	// 获取时间戳
	t1 := parse.Unix()
	// 获取纳秒级时间戳
	t2 := parse.UnixNano()
	t.Log(t1,t2)
}
