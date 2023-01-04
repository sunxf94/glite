package utils

import "time"

const (
	layoutDefault = "2006-01-02 15:04:05"
)

func Str2Time(s string) (time.Time, error) {
	return time.ParseInLocation(layoutDefault, s, time.Local)
}
