package store

import "github.com/dkeng/w4w/src/entity"

// LinkStore 链接存储接口
type LinkStore interface {
	// Exist 判断URL是否存在
	Exist(url string)
	// Add 添加链接
	Add(link *entity.Link)
	// QueryByShort 根据短链接获取
	QueryByShort(short string) *entity.Link
}
