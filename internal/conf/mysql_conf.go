package conf

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetDbSetting() (mysql.Config, logger.Interface) {
	// 设置时区为东八区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 数据库配置
	config := mysql.Config{
		User:   "root",
		Passwd: "hongjun500",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "mall_go",
		Loc:    loc,
		// time.Time 与 mysql 日期类型互转
		ParseTime: true,
		// 允许本地密码认证
		AllowNativePasswords: true,
	}
	// gorm 的日志配置
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	return config, gormLogger
}
