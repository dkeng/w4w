package mysql

import (
	"time"

	"github.com/dkeng/pkg/store"
	"github.com/dkeng/w4w/src/entity"
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
	var count int64
	r.Db.Model(&entity.RedirectRecord{}).Where("created_at >= ? AND created_at <= ?", startTime, endTime).Count(&count)
	return 0
}

// RankByStartTimeAndEndTime 获取今天访问排行榜
func (r *RedirectRecordStore) RankByStartTimeAndEndTime(startTime, endTime time.Time) []map[string]interface{} {
	var result []map[string]interface{}
	rows, err := r.Db.Raw(`SELECT link_id,COUNT(*) AS count FROM redirect_records  
	WHERE  created_at >= ? AND created_at <= ? 
	GROUP BY link_id LIMIT 0,10`, startTime, endTime).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result, err = store.ScanRowsToMaps(rows)
	if err != nil {
		return nil
	}
	return result
}
