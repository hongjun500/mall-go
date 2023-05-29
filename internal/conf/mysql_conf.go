package conf

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	Once sync.Once
	Db   *gorm.DB
	err  error
)

type DBConnector interface {
	Connect() (*gorm.DB, error)
}

type MySQLConnector struct {
	/*Once sync.Once
	Db   *gorm.DB
	err  error*/
}

func InitMySQLConn() (*gorm.DB, error) {
	// func (c *MySQLConnector) Connect() (*gorm.DB, error) {
	Once.Do(func() {

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
		sqldb, err := sql.Open("mysql", config.FormatDSN())

		Db, err = gorm.Open(gmysql.New(gmysql.Config{
			Conn: sqldb,
		}), &gorm.Config{})

		if err != nil {
			err = fmt.Errorf("MySQL Connected Fail, ERR = %v", err)
			return
		}
		fmt.Println("MySQL Connected!")
	})
	return Db, err
}
