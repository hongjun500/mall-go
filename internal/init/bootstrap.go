package init

import (
	"fmt"
	"github.com/hongjun500/mall-go/internal/sql-factory"
)

var (
	SqlSession *sql_factory.SqlFactory
)

// StartUp 启动初始化
func StartUp() {
	gormMySQL, _ := GormMySQL()
	SqlSession = sql_factory.NewSqlFactory(gormMySQL, nil)

	fmt.Println("gormMySQL = ", gormMySQL)
	fmt.Println("sqlFactory = ", SqlSession)
}
