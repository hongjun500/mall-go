package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/hongjun500/mall-go/internal/conf"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	once = sync.Once{}
)

func newMySQL(properties conf.DatabaseConfigProperties) (db *sql.DB, err error) {
	// 设置时区为东八区
	loc, _ := time.LoadLocation(conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties.Loc)
	config := mysql.Config{
		User:   properties.GormMysqlConfigProperties.Username,
		Passwd: properties.GormMysqlConfigProperties.Password,
		Net:    "tcp",
		Addr:   properties.GormMysqlConfigProperties.Host + ":" + properties.GormMysqlConfigProperties.Port,
		DBName: properties.GormMysqlConfigProperties.Database,
		Loc:    loc,
		// time.Time 与 mysql 日期类型互转
		ParseTime: conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties.ParseTime,
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
func NewGormMySQL(properties conf.DatabaseConfigProperties) (db *gorm.DB, err error) {
	// gorm 对于 MySQL 的连接

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
				SlowThreshold:             properties.GormMysqlConfigProperties.GormSlowThreshold,             // Slow SQL threshold
				LogLevel:                  logger.LogLevel(properties.GormMysqlConfigProperties.GormLogLevel), // Log level
				IgnoreRecordNotFoundError: properties.GormMysqlConfigProperties.GormIgnoreRecordNotFoundError, // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      properties.GormMysqlConfigProperties.GormParameterizedQueries,      // Don't include params in the SQL log
				Colorful:                  properties.GormMysqlConfigProperties.GormColorful,                  // Disable color
			},
		)

		db, err = gorm.Open(gormMysql.New(gormMysql.Config{
			Conn: mySQL,
		}), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			err = fmt.Errorf("GormMySQL Connected Fail, ERR = %v", err)
		}
		fmt.Println("GormMySQL Connected!")
	})
	return db, err
}
