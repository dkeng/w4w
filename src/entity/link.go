package entity

// Link 链接
type Link struct {
	BaseEntity
	URL   string `gorm:"not null"`
	Title string `gorm:"null"`
	Short string `gorm:"not null;unique_index"`
}
