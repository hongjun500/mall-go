package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"gorm.io/gorm"
)

type RedisCli struct {
}

type SqlFactory struct {
	DbMySQL *gorm.DB
	// todo 改写 redis
	DbRedis RedisCli
}

type GinEngine struct {
	GinEngine *gin.Engine
}

var (
	SqlSession *SqlFactory
)

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
}

// StartUp 启动初始化
func StartUp() {
	gormMySQL, _ := database.NewGormMySQL()
	SqlSession = NewSqlSessionFactory(gormMySQL, nil)
	// Gin 初始化
	ginEngine := NewGinEngine()
	ginEngine.GinEngine.Run(":8080")
}

type SqlSessionFactory interface {
	NewSqlSessionFactory(args ...any) *SqlFactory
}

func NewSqlSessionFactory(args ...any) *SqlFactory {
	factory := &SqlFactory{
		DbMySQL: nil,
		DbRedis: RedisCli{},
	}
	for _, arg := range args {
		switch val := arg.(type) {
		case *gorm.DB:
			factory.DbMySQL = val
		case RedisCli:
			factory.DbRedis = val
		}
	}
	return factory
}
