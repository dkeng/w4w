package entity

import "time"

// RequestRecord 请求记录
type RequestRecord struct {
	ID          int64         `gorm:"primary_key;unique_index"`
	StartTime   time.Time     `gorm:"not null;type:DATETIME"`
	EndTime     time.Time     `gorm:"not null;type:DATETIME"`
	Referer     string        `gorm:"null"`
	Path        string        `gorm:"null"`
	RawQuery    string        `gorm:"null"`
	Latency     time.Duration `gorm:"null"`
	ClientIP    string        `gorm:"null"`
	Method      string        `gorm:"null"`
	StatusCode  int           `gorm:"null"`
	UserAgent   string        `gorm:"null"`
	RequestSize int64         `gorm:"null"`
	WriterSize  int64         `gorm:"null"`
}
