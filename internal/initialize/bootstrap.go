package initialize

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hongjun500/mall-go/docs"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/routers"
	"github.com/hongjun500/mall-go/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartUpAdmin admin 模块启动初始化
func StartUpAdmin() *gin.Engine {

	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL(conf.GlobalDatabaseConfigProperties)

	// 拿到
	redisClient, _ := database.NewRedisClient(conf.GlobalDatabaseConfigProperties)

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, redisClient, nil)

	// 将与业务逻辑相关的封装到一个结构体中
	coreService := services.NewCoreService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := routers.NewCoreRouter(coreService)

	// 初始化 gin 引擎
	ginEngine := NewGinEngine().GinEngine

	// 初始化路由分组
	initGroupRouter(coreRouter, ginEngine)

	return ginEngine
}

// NewGinEngine 初始化 gin 引擎
func NewGinEngine() *gin_common.GinEngine {
	r := gin.Default()
	gin.SetMode(conf.GlobalAdminServerConfigProperties.GinRunMode)
	engine := &gin_common.GinEngine{GinEngine: r}
	// 强制日志颜色化
	// gin.ForceConsoleColor()
	// 限流中间件
	/*routers.go.Use(limits.RequestSizeLimiter(10))
	routers.go.Use(cors.Default())*/
	// r.Use(gin.Recovery())

	// 跨域中间件
	r.Use(mid.GinCORSMiddleware())
	return engine
}

// initGroupRouter 初始化路由分组
func initGroupRouter(coreRouter *routers.CoreRouter, ginEngine *gin.Engine) {
	// 必须要写上这一行很奇怪
	// 解释：必须要导入 swagger 的包，即 docs, 不然 swagger 无法生成文档
	// docs.SwaggerInfo.Version = "1.0"
	/*
		docs.SwaggerInfo.Title = "mall-go"
		docs.SwaggerInfo.Description = "mall-go 项目接口文档"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.BasePath = "/api/v1"*/

	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册多个路由
	coreRouter.GroupUmsAdminRouter(ginEngine)
	coreRouter.GroupUmsMenuRouter(ginEngine)
	coreRouter.GroupUmsResourceCategoryRouter(ginEngine)
	coreRouter.GroupUmsResourceRouter(ginEngine)
	coreRouter.GroupUmsRoleRouter(ginEngine)
	coreRouter.GroupUmsMemberLevelRouter(ginEngine)
}

// StartUpPortal portal 模块启动初始化
func StartUpPortal() *gin.Engine {
	// todo portal 的初始化
	return gin.New()
}

// StartUpSearch search 模块启动初始化
func StartUpSearch() *gin.Engine {
	// todo search 的初始化
	return gin.New()
}
