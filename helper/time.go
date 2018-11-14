package helper

import "time"

// StrToTime 字符串转时间戳
func StrToTime(date string) int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", date)
	return t.Unix()
}

// NowTime 获取当前时间戳
func NowTime() int64 {
	t := time.Now()
	return t.Unix()
}
