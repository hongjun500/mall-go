package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Es struct {
	// todo
}

type DbFactory struct {
	// 基于 gorm 的 MySQL 连接
	GormMySQL *gorm.DB

	// 基于 go-redis 的 Redis 连接
	RedisCli *redis.Client
	// todo 改写 es
	Es *Es
}

func NewDbFactory(args ...any) *DbFactory {
	factory := &DbFactory{
		GormMySQL: nil,
		RedisCli:  nil,
		Es:        nil,
	}
	for _, arg := range args {
		switch val := arg.(type) {
		case *gorm.DB:
			factory.GormMySQL = val
		case *redis.Client:
			factory.RedisCli = val
		case Es:
			// todo 改写 es
			factory.Es = &val
		}
	}
	return factory
}
