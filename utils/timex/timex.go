package timex

import (
	"fmt"
	"strings"
	"time"
)

const (
	defaultLayout = "2006-01-02 15:04:05"
)

// 耗时统计
func TimeConsuming() func() {
	start := time.Now()

	return func() {
		fmt.Println(fmt.Sprintf("耗时: %0.3fs", time.Now().Sub(start).Seconds()))
	}
}

func Format(t time.Time, layout string) string {
	return t.Format(format2layout(layout))
}

func FormatTime(t time.Time) string {
	return t.Format(defaultLayout)
}

func FormatYMD(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatMD(t time.Time) string {
	return t.Format("01月02日")
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse(defaultLayout, s)
}

func ParseYMD(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// 获取该月的第一天初始时间
func FirstMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// 获取该月的第一天初始时间戳
func FirstMonthUnix(t time.Time) int64 {
	return FirstMonth(t).Unix()
}

// 获取该月的最后一天时间
func LastMonth(t time.Time) time.Time {
	nextMonth := t.Month() + 1
	next := time.Date(t.Year(), nextMonth, 1, 0, 0, 0, 0, t.Location())
	next = next.Add(time.Microsecond * -1)
	return next
}

// 获取该月的最后一天时间戳
func LastMonthUnix(t time.Time) int64 {
	return LastMonth(t).Unix()
}

// format转layout
func format2layout(format string) string {
	format = strings.Trim(format, " ")
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "M", "01", 1)
	layout = strings.Replace(layout, "D", "02", 1)
	layout = strings.Replace(layout, "h", "15", 1)
	layout = strings.Replace(layout, "m", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	return layout
}
