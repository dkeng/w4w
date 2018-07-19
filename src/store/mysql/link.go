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

// Exist 判断URL是否存在
func (l *LinkStore) Exist(url string) {

}

// Add 添加链接
func (l *LinkStore) Add(link *entity.Link) {

}

// QueryByShort 根据短链接获取
func (l *LinkStore) QueryByShort(short string) *entity.Link {
	return nil
}
