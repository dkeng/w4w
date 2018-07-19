package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
)

// RedirectRecordStore 链接储存
type RedirectRecordStore struct {
	baseMySQLStore
}

// Init 初始化
func (r *RedirectRecordStore) Init(db *gorm.DB) *RedirectRecordStore {
	r.Db = db
	r.Name = "RedirectRecordStore"
	return r
}

// CountByStartTimeAndEndTime 查询指定时间范围内的数量
func (r *RedirectRecordStore) CountByStartTimeAndEndTime(startTime, endTime time.Time) int64 {
	return 0
}
