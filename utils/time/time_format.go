package time

import "time"

// 时间格式化
func TimeFormat() string {
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	return nowTime
}
