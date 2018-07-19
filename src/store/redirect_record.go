package store

import "time"

// RedirectRecordStore 重定向记录存储接口
type RedirectRecordStore interface {
	// CountByStartTimeAndEndTime 查询指定时间范围内的数量
	CountByStartTimeAndEndTime(startTime, endTime time.Time) int64
}
