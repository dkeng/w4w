package custom

// LinkRank 排行
type LinkRank struct {
	// LinkID 链接ID
	LinkID int64 `gorm:"column:link_id"`
	// 访问数量
	Count int64 `gorm:"column:count"`
}
