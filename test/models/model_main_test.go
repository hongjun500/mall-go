package models

import (
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"gorm.io/gorm"
	"testing"
)

var TestModelGormMySQL *gorm.DB

func TestMain(m *testing.M) {
	// 通过 gorm 拿到 MySQL 数据库连接
	conf.InitConfigProperties()

	TestModelGormMySQL, _ = database.NewGormMySQL(conf.GlobalDatabaseConfigProperties)
	m.Run()
}
