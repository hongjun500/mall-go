package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hongjun500/mall-go/internal/conf"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newMySQL(properties conf.GormMysqlConfigProperties) (db *sql.DB, err error) {
	// 设置时区为东八区
	loc, _ := time.LoadLocation(properties.Loc)
	config := mysql.Config{
		User:   properties.Username,
		Passwd: properties.Password,
		Net:    "tcp",
		Addr:   properties.Host + ":" + properties.Port,
		DBName: properties.Database,
		Loc:    loc,
		// time.Time 与 mysql 日期类型互转
		ParseTime: properties.ParseTime,
		// 允许本地密码认证
		AllowNativePasswords: true,
	}

	// mysql 连接
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("MySQL Connected Fail, ERR = %v", err)
	}
	err = db.Ping()
	if err != nil {
		// 关闭连接池
		_ = db.Close()
		return nil, fmt.Errorf("MySQL Connection Failed: %v", err)
	}

	fmt.Println("MySQL Connected!")
	return db, err
}

// NewGormMySQL 初始化 gorm 对于 MySQL 的连接
func NewGormMySQL(properties conf.GormMysqlConfigProperties) (*gorm.DB, error) {
	// gorm 对于 MySQL 的连接
	var db *gorm.DB
	var err error
	once := sync.Once{}
	once.Do(func() {
		mySQL, err := newMySQL(properties)
		if err != nil {
			log.Fatalln("NewGormMySQL Fail, ERR = ", err)
			return
		}

		// gorm 的日志配置
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             properties.GormSlowThreshold,             // Slow SQL threshold
				LogLevel:                  logger.LogLevel(properties.GormLogLevel), // Log level
				IgnoreRecordNotFoundError: properties.GormIgnoreRecordNotFoundError, // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      properties.GormParameterizedQueries,      // Don't include params in the SQL log
				Colorful:                  properties.GormColorful,                  // Disable color
			},
		)

		db, err = gorm.Open(gormMysql.New(gormMysql.Config{
			Conn: mySQL,
		}), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			log.Fatalln("GormMySQL Connected Fail, ERR = ", err)
		}
	})
	fmt.Println("GormMySQL Connected!")
	return db, err
}
