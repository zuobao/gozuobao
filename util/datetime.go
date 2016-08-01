package util

import "time"

// 当前毫秒数
func NowMs() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetMs(t *time.Time) int64 {
	return t.UnixNano() / 1e6
}

func TimestampToTime(ts int64) time.Time {
	return time.Unix(ts/1000, (ts%1000)*int64(time.Millisecond))
}

func AddDate(ts int64, years, months, days int) int64 {
	t := TimestampToTime(ts)
	newT := t.AddDate(years, months, days)
	return GetMs(&newT)
}

func GetDatePart(t *time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}