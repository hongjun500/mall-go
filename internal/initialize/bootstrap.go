package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/services"
)

type GinEngine struct {
	GinEngine *gin.Engine
}

func NewGinEngine() *GinEngine {
	r := gin.Default()
	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 路由分组
	api := r.Group("/api")
	{
		// 用户注册
		api.POST("/ums/admin/register", nil)
	}

	engine := &GinEngine{GinEngine: r}

	return engine

	/*// Gin 初始化
	ginEngine := NewGinEngine()
	ginEngine.GinEngine.Run(":8080")*/
}

// StartUp 启动初始化
func StartUp() {
	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL()

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, nil)

	// 将与业务逻辑相关的封装到一个结构体中
	_ = services.InitCoreService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中

}
