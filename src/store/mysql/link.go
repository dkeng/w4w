package mysql

import (
	"github.com/dkeng/w4w/src/entity"
	"github.com/jinzhu/gorm"
)

// LinkStore 链接储存
type LinkStore struct {
	baseMySQLStore
}

// Init 初始化
func (l *LinkStore) Init(db *gorm.DB) *LinkStore {
	l.Db = db
	l.Name = "LinkStore"
	return l
}

// ExistURL 判断URL是否存在
func (l *LinkStore) ExistURL(url string) bool {
	var count int64
	l.Db.Model(&entity.Link{}).Where("url = ?", url).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// ExistShort 判断短链接是否存在
func (l *LinkStore) ExistShort(url string) bool {
	var count int64
	l.Db.Model(&entity.Link{}).Where("short = ?", url).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// Add 添加链接
func (l *LinkStore) Add(link *entity.Link) bool {
	if err := l.Db.Create(&link).Error; err != nil {
		return false
	}
	return true
}

// QueryByShort 根据短链接获取
func (l *LinkStore) QueryByShort(short string) *entity.Link {
	var link entity.Link
	err := l.Db.First(&link, "short = ?", short).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		return nil
	}
	return &link
}
