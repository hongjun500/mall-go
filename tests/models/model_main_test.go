package models

import (
	"testing"

	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"gorm.io/gorm"
)

var TestModelGormMySQL *gorm.DB
var TestModelDbFactory *database.DbFactory

func TestMain(m *testing.M) {
	// 通过 gorm 拿到 MySQL 数据库连接
	conf.InitAdminConfigProperties()
	TestModelGormMySQL, _ = database.NewGormMySQL(conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties)
	TestModelDbFactory = database.NewDbFactory(TestModelGormMySQL)

	m.Run()
}
