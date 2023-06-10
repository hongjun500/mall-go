/**
 * @author hongjun500
 * @date 2023/6/8
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 所有中间组件需要用到的变量都放这里定义并且完成初始化
 */

package common

import (
	"context"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"testing"
)

// 当前包下的全局变量
var (
	client *redis.Client
	ctx    = context.Background()
	db     *gorm.DB
)

func TestMain(m *testing.M) {
	conf.InitAdminConfigProperties()
	var err error
	client, err = database.NewRedisClient(conf.GlobalDatabaseConfigProperties)
	db, err = database.NewGormMySQL(conf.GlobalDatabaseConfigProperties)
	if err != nil {
		return
	}
	log.Print("gorm db: ", db)
	if err != nil {
		return // todo
	}
	log.Print("redis client: ", client)
	m.Run()
}
