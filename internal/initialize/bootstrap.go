package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/gin_common/security"
	"github.com/hongjun500/mall-go/internal/routers/r_mall_admin"
	"github.com/hongjun500/mall-go/internal/routers/r_mall_search"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_search"
)

// StartUpAdmin admin 模块启动初始化
func StartUpAdmin() *gin.Engine {

	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL(conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties)

	// 拿到 Redis 数据库连接
	redisClient, _ := database.NewRedisClient(conf.GlobalDatabaseConfigProperties.RedisConfigProperties)

	// 将与数据库相关的封装到一个结构体中
	sqlSessionFactory := database.NewDbFactory(gormMySQL, redisClient, nil)

	security.SetDbFactory(sqlSessionFactory)

	// 初始化 gin 引擎
	ginEngine := NewAdminGinEngine()

	// 将与业务逻辑相关的封装到一个结构体中
	coreService := s_mall_admin.NewCoreAdminService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := r_mall_admin.NewCoreAdminRouter(coreService)

	// 初始化路由分组
	r_mall_admin.InitAdminGroupRouter(coreRouter, ginEngine)

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
	// 全局错误处理中间件
	ginEngine.Use(mid.ErrorMiddleware())
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
	// 通过 gorm 拿到 MySQL 数据库连接
	gormMySQL, _ := database.NewGormMySQL(conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties)
	// 拿到 elasticsearch 的连接
	es, _ := database.NewEsTypedClient(conf.GlobalDatabaseConfigProperties.ElasticSearchConfigProperties)

	sqlSessionFactory := database.NewDbFactory(gormMySQL, nil, es)

	coreSearchService := s_mall_search.NewCoreSearchService(sqlSessionFactory)

	// 将与路由相关的封装到一个结构体中
	coreRouter := r_mall_search.NewCoreSearchRouter(coreSearchService)

	// 初始化 gin 引擎
	ginEngine := NewSearchGinEngine()

	// 初始化路由分组
	r_mall_search.InitSearchGroupRouter(coreRouter, ginEngine)
	return ginEngine
}

// NewSearchGinEngine 初始化 gin 引擎
func NewSearchGinEngine() *gin.Engine {
	ginEngine := gin.Default()

	gin.SetMode(conf.GlobalSearchServerConfigProperties.GinRunMode)

	// 强制日志颜色化
	// gin.ForceConsoleColor()
	// 限流中间件
	/*r.Use(limits.RequestSizeLimiter(10))
	r.Use(cors.Default())*/

	// 跨域中间件
	ginEngine.Use(mid.GinCORSMiddleware())
	// 全局错误处理中间件
	ginEngine.Use(mid.ErrorMiddleware())
	return ginEngine
}
