package entity

import "time"

// RedirectRecord 重定向记录
type RedirectRecord struct {
	ID          int64     `gorm:"primary_key;unique_index"`
	LinkID      int64     `gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;type:DATETIME"`
	Referer     string    `gorm:"null"`
	Path        string    `gorm:"null"`
	RawQuery    string    `gorm:"null"`
	ClientIP    string    `gorm:"null"`
	Method      string    `gorm:"null"`
	StatusCode  int       `gorm:"null"`
	UserAgent   string    `gorm:"null"`
	RequestSize int64     `gorm:"null"`
	WriterSize  int64     `gorm:"null"`
}
