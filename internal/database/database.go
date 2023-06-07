package database

import (
	"gorm.io/gorm"
)

type Redis struct {
	// todo
}
type Es struct {
	// todo
}

type DbFactory struct {
	// 基于 gorm 的 MySQL 连接
	GormMySQL *gorm.DB

	// todo 改写 redis
	Redis *Redis
	// todo 改写 es
	Es *Es
}

func NewDbFactory(args ...any) *DbFactory {
	factory := &DbFactory{
		GormMySQL: nil,
		Redis:     nil,
		Es:        nil,
	}
	for _, arg := range args {
		switch val := arg.(type) {
		case *gorm.DB:
			factory.GormMySQL = val
		case Redis:
			// todo 改写 redis
			factory.Redis = &val
		case Es:
			// todo 改写 es
			factory.Es = &val
		}
	}
	return factory
}
