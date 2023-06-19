package initialize

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hongjun500/mall-go/docs"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/gin_common/security"
	"github.com/hongjun500/mall-go/internal/routers"
	"github.com/hongjun500/mall-go/internal/services"
)

// StartUpAdmin admin 模块启动初始化
func StartUpAdmin() *gin.Engine {

	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL(conf.GlobalDatabaseConfigProperties)

	// 拿到 Redis 数据库连接
	redisClient, _ := database.NewRedisClient(conf.GlobalDatabaseConfigProperties)

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, redisClient, nil)

	security.SetDbFactory(sqlSessionFactory)

	// 初始化 gin 引擎
	ginEngine := NewAdminGinEngine()

	// 将与业务逻辑相关的封装到一个结构体中
	coreService := services.NewCoreAdminService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := routers.NewCoreAdminRouter(coreService)

	// 初始化路由分组
	routers.InitAdminGroupRouter(coreRouter, ginEngine)

	return ginEngine
}

// NewAdminGinEngine 初始化 gin 引擎
func NewAdminGinEngine() *gin.Engine {
	ginEngine := gin.Default()

	gin.SetMode(conf.GlobalAdminServerConfigProperties.GinRunMode)

	// 强制日志颜色化
	// gin.ForceConsoleColor()
	// 限流中间件
	/*r.Use(limits.RequestSizeLimiter(10))
	r.Use(cors.Default())*/

	// 跨域中间件
	ginEngine.Use(mid.GinCORSMiddleware())
	return ginEngine
}

// StartUpPortal portal 模块启动初始化
func StartUpPortal() *gin.Engine {
	// todo portal 的初始化
	engine := gin.New()
	return engine
}

// StartUpSearch search 模块启动初始化
func StartUpSearch() *gin.Engine {
	// todo search 的初始化
	engine := gin.New()

	// 拿到 es 的连接

	sqlSessionFactory := database.NewDbFactory(nil, nil, "es")

	coreSearchService := services.NewCoreSearchService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := routers.NewCoreSearchRouter(coreSearchService)

	// 初始化路由分组
	routers.InitSearchGroupRouter(coreRouter, engine)
	return engine
}
