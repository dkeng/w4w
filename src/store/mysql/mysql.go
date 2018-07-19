package mysql

import (
	"github.com/jinzhu/gorm"
)

// BaseMySQLStore shared DB data
type baseMySQLStore struct {
	Db   *gorm.DB
	Name string
}
