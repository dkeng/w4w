package store

import "time"

// RedirectRecordStore 重定向记录存储接口
type RedirectRecordStore interface {
	// CountByStartTimeAndEndTime 查询指定时间范围内的数量
	CountByStartTimeAndEndTime(startTime, endTime time.Time) int64
	// 获取访问前100 获取访问前100
	RankTop100() []map[string]interface{}
	// RankByStartTimeAndEndTime 获取今天访问排行榜
	RankByStartTimeAndEndTime(startTime, endTime time.Time) []map[string]interface{}
}
