package mysql

import (
	"time"

	"github.com/dkeng/w4w/src/entity"
	"github.com/dkeng/w4w/src/entity/custom"
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

// RankTop100 获取访问前100
func (r *RedirectRecordStore) RankTop100() []*custom.LinkRank {
	var linkRanks []*custom.LinkRank
	rows, err := r.Db.Raw(` SELECT link_id,count FROM (
		SELECT link_id,COUNT(*) AS count FROM redirect_records  
		GROUP BY link_id 
		) temp ORDER BY count DESC     LIMIT 0,100`).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var linkRank custom.LinkRank
		err := r.Db.ScanRows(rows, &linkRank)
		if err != nil {
			return nil
		}
		linkRanks = append(linkRanks, &linkRank)
	}
	return linkRanks
}

// RankByStartTimeAndEndTime 获取今天访问排行榜
func (r *RedirectRecordStore) RankByStartTimeAndEndTime(startTime, endTime time.Time) []*custom.LinkRank {
	var linkRanks []*custom.LinkRank
	rows, err := r.Db.Raw(`SELECT link_id,count FROM (
		SELECT link_id,COUNT(*) AS count FROM redirect_records  
		WHERE  created_at >= ? AND created_at <= ?
		GROUP BY link_id 
		) temp ORDER BY count DESC     LIMIT 0,10`, startTime, endTime).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var linkRank custom.LinkRank
		err := r.Db.ScanRows(rows, &linkRank)
		if err != nil {
			return nil
		}
		linkRanks = append(linkRanks, &linkRank)
	}
	return linkRanks
}
