package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/routers"
	"github.com/hongjun500/mall-go/internal/services"
)

// StartUp 启动初始化
func StartUp() *gin.Engine {
	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL()

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, nil)

	// 将与业务逻辑相关的封装到一个结构体中
	coreService := services.InitCoreService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := routers.InitCoreRouter(coreService)

	// 初始化 gin 引擎
	ginEngine := InitGinEngine().GinEngine

	// 初始化路由分组
	InitGroupRouter(coreRouter, ginEngine)

	return ginEngine
}

// InitGinEngine 初始化 gin 引擎
func InitGinEngine() *gin_common.GinEngine {
	r := gin.Default()
	// 强制日志颜色化
	gin.ForceConsoleColor()

	engine := &gin_common.GinEngine{GinEngine: r}
	// 强制日志颜色化
	gin.ForceConsoleColor()
	// 日志中间件
	r.Use(gin.Logger())
	// 限流中间件
	/*routers.go.Use(limits.RequestSizeLimiter(10))
	routers.go.Use(cors.Default())*/
	// r.Use(gin.Recovery())
	// gin.SetMode(gin)
	return engine
}

// InitGroupRouter 初始化路由分组
func InitGroupRouter(coreRouter *routers.CoreRouter, ginEngine *gin.Engine) {
	coreRouter.GroupUmsAdminRouter(ginEngine)
}
