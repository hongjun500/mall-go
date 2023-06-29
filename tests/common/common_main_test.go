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
	"log"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 当前包下的全局变量
var (
	dbFactory *database.DbFactory
	redisCli  *redis.Client
	typedCli  *elasticsearch.TypedClient
	esCli     *elasticsearch.Client
	ctx       = context.Background()
	db        *gorm.DB
)

func TestMain(m *testing.M) {
	conf.InitAdminConfigProperties()
	var err error
	redisCli, err = database.NewRedisClient(conf.GlobalDatabaseConfigProperties.RedisConfigProperties)
	db, err = database.NewGormMySQL(conf.GlobalDatabaseConfigProperties.GormMysqlConfigProperties)
	es, err := database.NewEsTypedClient(conf.GlobalDatabaseConfigProperties.ElasticSearchConfigProperties)
	dbFactory = database.NewDbFactory(redisCli, db, es)
	if err != nil {
		return
	}
	typedCli = es.TypedCli
	esCli = es.Cli
	log.Print("gorm db: ", db)
	log.Print("redis client: ", redisCli)
	m.Run()
}
