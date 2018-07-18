package store

import (
	"github.com/dkeng/pkg/logger"
	"github.com/dkeng/w4w/src/entity"
	"github.com/spf13/viper"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Store 存储
type Store struct {
	DB *gorm.DB
}

// Open 打开存储
func (s *Store) Open() (err error) {
	// 初始化数据库
	dialect := viper.GetString("db.dialect")
	db, err := gorm.Open(dialect, viper.GetString("db.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 %s 连接失败: %s \n", dialect,
			errors.Wrap(err, "打开连接失败"),
		)
	}

	if db.DB().Ping() != nil {
		logger.Fatalf(
			"初始化 %s 连接失败: %s \n", dialect,
			errors.Wrap(err, "Ping失败"),
		)
	}

	db.LogMode(viper.GetBool("db.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("db.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("db.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(
		&entity.Link{},
		&entity.RedirectRecord{},
		&entity.RequestRecord{},
	)
	s.DB = db
	return
}

// Close 关闭存储
func (s *Store) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

// AllStore 所有存储
type AllStore struct {
}

// Init 初始化
func (m *AllStore) Init(s *Store) {

}

func getLimitOffset(page, perPage *int) *int {
	if *page <= 0 {
		*page = 1
	}
	if *perPage <= 0 {
		*perPage = 10
	}
	offset := (*page - 1) * *perPage
	return &offset
}

func switchDB(tran *gorm.DB, db *gorm.DB) *gorm.DB {
	if tran != nil {
		return tran
	}
	if db != nil {
		return db
	}
	panic(errors.New("事务切换失败"))
}
