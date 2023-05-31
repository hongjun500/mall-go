package sql_factory

import "gorm.io/gorm"

type RedisCli struct {
}

type SqlFactory struct {
	DbMySQL *gorm.DB
	// todo 改写 redis
	DbRedis RedisCli
}

func NewSqlFactory(args ...any) *SqlFactory {
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
