package conf

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

// InitConfigProperties 初始化配置属性
func InitConfigProperties() {

	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("../configs/")
	viper.AddConfigPath("../../configs/")
	viper.AddConfigPath("../../../configs/")
	// viper.AddConfigPath("../../../../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		// panic(err)
		log.Println("读取配置文件失败", err)

		// 使用默认配置
		log.Println("开始使用默认配置")
		initDefaultConfigProperties()
		log.Println("默认配置初始化完成")
	} else {
		log.Println("读取配置文件成功")
		// 监听配置项的修改
		viper.WatchConfig()
		var gorMysqlConfigProperties GormMysqlConfigProperties
		var redisConfigProperties RedisConfigProperties
		_ = viper.UnmarshalKey("server", &GlobalServerConfigProperties)
		_ = viper.UnmarshalKey("jwt", &GlobalJwtConfigProperties)
		_ = viper.UnmarshalKey("database.gorm_mysql", &gorMysqlConfigProperties)
		_ = viper.UnmarshalKey("database.redis", &redisConfigProperties)
		GlobalDatabaseConfigProperties = DatabaseConfigProperties{
			GormMysqlConfigProperties: gorMysqlConfigProperties,
			RedisConfigProperties:     redisConfigProperties,
		}
		log.Println("配置项初始化完成")
	}

}

// 使用默认配置项
func initDefaultConfigProperties() {
	GlobalServerConfigProperties = ServerConfigProperties{
		GinRunMode:  "debug",
		Host:        "localhost",
		Port:        "8080",
		ReadTimeout: 60,
	}
	GlobalJwtConfigProperties = JwtConfigProperties{
		TokenHeader: "Authorization",
		TokenHead:   "Bearer ",
		Expiration:  60 * 60 * 24 * 7,
		Secret:      "mall-go-jwt-secret",
	}
	GlobalDatabaseConfigProperties = DatabaseConfigProperties{
		GormMysqlConfigProperties: GormMysqlConfigProperties{
			Host:     "localhost",
			Port:     "3306",
			Database: "mall_go",
			Username: "root",
			Password: "root",
			Charset:  "utf8mb4_general_ci",
			// 时区
			Loc:                           "Asia/Shanghai",
			ParseTime:                     true,
			GormSlowThreshold:             200 * time.Millisecond,
			GormColorful:                  false,
			GormIgnoreRecordNotFoundError: true,
			GormParameterizedQueries:      true,
			// Gorm 日志级别
			// 1 Silent，2 Error，3 Warn，4 Info
			GormLogLevel: 3,
		},
		RedisConfigProperties: RedisConfigProperties{
			Host:     "localhost",
			Port:     "6379",
			Password: "",
			Database: 0,
		},
	}
}