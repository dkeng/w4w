package store

import "github.com/dkeng/w4w/src/entity"

// LinkStore 链接存储接口
type LinkStore interface {
	// ExistURL 判断URL是否存在
	ExistURL(url string) bool
	// ExistShort 判断短链接是否存在
	ExistShort(short string) bool
	// Add 添加链接
	Add(link *entity.Link) bool
	// QueryByShort 根据短链接获取
	QueryByShort(short string) *entity.Link
}
