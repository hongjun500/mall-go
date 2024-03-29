package conf

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	MYSQL_HOST            string
	REDIS_HOST            string
	ELASTICSEARCH_ADDRESS string
)

// init 初始化所有配置属性
func init() {

	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	REDIS_HOST = os.Getenv("REDIS_HOST")
	ELASTICSEARCH_ADDRESS = os.Getenv("ELASTICSEARCH_ADDRESS")

	InitAdminConfigProperties()
	InitPortalConfigProperties()
	InitSearchConfigProperties()
}

// InitAdminConfigProperties 初始化配置属性
func InitAdminConfigProperties() {
	log.Println("------ InitAdminConfigProperties ------")
	viper.SetConfigType("yml")
	viper.SetConfigName("admin-config")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("../configs/")
	viper.AddConfigPath("../../configs/")
	viper.AddConfigPath("../../../configs/")
	// viper.AddConfigPath("../../../../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		// viper.BindEnv("server.port", "SERVER_PORT")

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
		_ = viper.UnmarshalKey("server", &GlobalAdminServerConfigProperties)
		_ = viper.UnmarshalKey("jwt", &GlobalJwtConfigProperties)
		_ = viper.UnmarshalKey("database.gorm_mysql", &gorMysqlConfigProperties)
		_ = viper.UnmarshalKey("database.redis", &redisConfigProperties)
		if strings.HasPrefix(gorMysqlConfigProperties.Host, "$") {
			gorMysqlConfigProperties.Host = MYSQL_HOST
		}
		if strings.HasPrefix(redisConfigProperties.Host, "$") {
			redisConfigProperties.Host = REDIS_HOST
		}
		GlobalDatabaseConfigProperties = DatabaseConfigProperties{
			GormMysqlConfigProperties: gorMysqlConfigProperties,
			RedisConfigProperties:     redisConfigProperties,
		}
		log.Println("配置项初始化完成")
	}

}
func InitPortalConfigProperties() {
	// todo 改为 portal 服务所需配置项
	log.Println("------ InitPortalConfigProperties ------")
	viper.SetConfigType("yml")
	viper.SetConfigName("portal-config")
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
		// initDefaultConfigProperties()
		log.Println("默认配置初始化完成")
	} else {
		log.Println("读取配置文件成功")
		// 监听配置项的修改
		viper.WatchConfig()
		// var gorMysqlConfigProperties GormMysqlConfigProperties
		// var redisConfigProperties RedisConfigProperties
		_ = viper.UnmarshalKey("server", &GlobalPortalServerConfigProperties)
		// _ = viper.UnmarshalKey("database.gorm_mysql", &gorMysqlConfigProperties)
		// _ = viper.UnmarshalKey("database.redis", &redisConfigProperties)
		// GlobalDatabaseConfigProperties = DatabaseConfigProperties{
		// 	GormMysqlConfigProperties: gorMysqlConfigProperties,
		// 	RedisConfigProperties:     redisConfigProperties,
	}
	log.Println("配置项初始化完成")
}

func InitSearchConfigProperties() {
	log.Println("------ InitSearchConfigProperties ------")
	viper.SetConfigType("yml")
	viper.SetConfigName("search-config")
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
		// initDefaultConfigProperties()
		log.Println("默认配置初始化完成")
	} else {
		log.Println("读取配置文件成功")
		// 监听配置项的修改
		viper.WatchConfig()
		// var gorMysqlConfigProperties GormMysqlConfigProperties
		// var redisConfigProperties RedisConfigProperties
		var esConfigProperties ElasticSearchConfigProperties
		_ = viper.UnmarshalKey("server", &GlobalSearchServerConfigProperties)
		// _ = viper.UnmarshalKey("database.gorm_mysql", &gorMysqlConfigProperties)
		// _ = viper.UnmarshalKey("database.redis", &redisConfigProperties)
		_ = viper.UnmarshalKey("database.elasticsearch", &esConfigProperties)
		if strings.HasPrefix(esConfigProperties.Addresses[0], "$") {
			esConfigProperties.Addresses = strings.Split(ELASTICSEARCH_ADDRESS, ",")
		}
		GlobalDatabaseConfigProperties.ElasticSearchConfigProperties = esConfigProperties
		log.Println("配置项初始化完成")
	}

}

func setDBEnv(str string) {
	if strings.HasPrefix(str, "$") {
		str = MYSQL_HOST
	}

	// 设置环境变量
	_ = os.Setenv("SERVER_PORT", "8080")
	_ = os.Setenv("MYSQL_HOST", "32")
	_ = os.Setenv("MYSQL_PORT", "3306")
}

// 使用默认配置项
func initDefaultConfigProperties() {
	GlobalAdminServerConfigProperties = ServerConfigProperties{
		ApplicationName: "mall_admin",
		GinRunMode:      "debug",
		Port:            "8080",
		ReadTimeout:     60,
	}
	GlobalPortalServerConfigProperties = ServerConfigProperties{
		ApplicationName: "mall-portal",
		GinRunMode:      "debug",
		Port:            "8081",
		ReadTimeout:     60,
	}
	GlobalSearchServerConfigProperties = ServerConfigProperties{
		ApplicationName: "mall_search",
		GinRunMode:      "debug",
		Port:            "8082",
		ReadTimeout:     60,
	}
	GlobalJwtConfigProperties = JwtConfigProperties{
		TokenHeader: "Authorization",
		TokenHead:   "Bearer ",
		Expiration:  60 * 60 * 24 * 7,
		Secret:      "mall-go-security-secret",
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
		ElasticSearchConfigProperties: ElasticSearchConfigProperties{
			Addresses:  []string{"https://localhost:9200"},
			Username:   "elastic",
			Password:   "elastic",
			CACertPath: "./configs/http_ca.crt",
		},
	}
}
