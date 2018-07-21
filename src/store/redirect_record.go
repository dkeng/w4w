package store

import (
	"time"

	"github.com/dkeng/w4w/src/entity/custom"
)

// RedirectRecordStore 重定向记录存储接口
type RedirectRecordStore interface {
	// CountByStartTimeAndEndTime 查询指定时间范围内的数量
	CountByStartTimeAndEndTime(startTime, endTime time.Time) int64
	// 获取访问前100 获取访问前100
	RankTop100() []*custom.LinkRank
	// RankByStartTimeAndEndTime 获取今天访问排行榜
	RankByStartTimeAndEndTime(startTime, endTime time.Time) []*custom.LinkRank
}
