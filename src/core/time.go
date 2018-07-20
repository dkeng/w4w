package core

import (
	"time"
)

// TodayStartEndTime 今天开始时间和结束时间
func TodayStartEndTime() (time.Time, time.Time) {
	today := time.Now().Format("2006-01-02")
	layout := "2006-01-02 15:04:05"
	start, _ := time.Parse(layout, today+" 00:00:00")
	end, _ := time.Parse(layout, today+" 23:59:59")
	return start, end
}
