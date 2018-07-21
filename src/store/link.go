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
	// QueryIDByShort 根据短链接获取ID
	QueryIDByShort(short string) int64
	// QueryInID 根据多个ID获取多个内容
	QueryInID(ids ...int64) []*entity.Link
	// UpdateTitleByURL 根据URL修改标题
	UpdateTitleByURL(url, title string)
}
