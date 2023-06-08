package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/docs"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/routers"
	"github.com/hongjun500/mall-go/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartUp 启动初始化
func StartUp() *gin.Engine {
	// 初始化配置文件的属性
	conf.InitConfigProperties()

	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL(conf.GlobalDatabaseConfigProperties)

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, nil, nil)

	// 将与业务逻辑相关的封装到一个结构体中
	coreService := services.NewCoreService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := routers.NewCoreRouter(coreService)

	// 初始化 gin 引擎
	ginEngine := NewGinEngine().GinEngine

	// 初始化路由分组
	InitGroupRouter(coreRouter, ginEngine)

	return ginEngine
}

// NewGinEngine 初始化 gin 引擎
func NewGinEngine() *gin_common.GinEngine {
	r := gin.Default()
	gin.SetMode(conf.GlobalServerConfigProperties.GinRunMode)
	engine := &gin_common.GinEngine{GinEngine: r}
	// 强制日志颜色化
	// gin.ForceConsoleColor()
	// 限流中间件
	/*routers.go.Use(limits.RequestSizeLimiter(10))
	routers.go.Use(cors.Default())*/
	// r.Use(gin.Recovery())
	// gin.SetMode(gin)
	return engine
}

// InitGroupRouter 初始化路由分组
func InitGroupRouter(coreRouter *routers.CoreRouter, ginEngine *gin.Engine) {
	// 必须要写上这一行很奇怪
	docs.SwaggerInfo.Version = "1.0"

	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 多个路由
	coreRouter.GroupUmsAdminRouter(ginEngine)
}
