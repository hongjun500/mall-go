package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/redis/go-redis/v9"
)

// NewRedisClient 初始化 redis 连接
func NewRedisClient(properties conf.DatabaseConfigProperties) (*redis.Client, error) {
	var client *redis.Client
	var err error
	once := sync.Once{}
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     properties.RedisConfigProperties.Host + ":" + properties.RedisConfigProperties.Port,
			Password: properties.RedisConfigProperties.Password,
			DB:       properties.RedisConfigProperties.Database,
		})
		ctx := context.Background()
		ping := client.Ping(ctx)
		if ping.Err() != nil {
			client = nil
			err = fmt.Errorf("redis Connected Fail, ERR = %v", ping.Err())
		}
	})

	if err != nil || client == nil {
		log.Fatalln("NewRedisClient Fail, ERR = ", err)
	}
	fmt.Println("Redis Connected!")
	return client, nil
}
