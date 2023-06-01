package database

import (
	"database/sql"
	"fmt"
	"github.com/hongjun500/mall-go/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type DbFactory struct {
	// 基于 gorm 的 MySQL 连接
	GormMySQL *gorm.DB

	// todo 改写 redis
	Redis string
	// todo 改写 es
	Es string
}

type DbSessionFactory interface {
	NewDbSessionFactory(args ...any) *DbFactory
}

// NewGormMySQL 初始化 gorm 对于 MySQL 的连接
func NewGormMySQL() (db *gorm.DB, err error) {
	// 拿到配置
	config, gormLogger := conf.GetDbSetting()
	// 初始化 gorm 对于 MySQL 的连接
	var Once sync.Once
	Once.Do(func() {
		// todo 这里可以通过别的方式获取配置
		sqlDb, err := sql.Open("mysql", config.FormatDSN())
		db, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDb,
		}), &gorm.Config{
			Logger: gormLogger,
		})

		if err != nil {
			err = fmt.Errorf("MySQL Connected Fail, ERR = %v", err)
			return
		}
		fmt.Println("MySQL Connected!")
	})
	return db, err
}
