package conf

import "time"

// 全局配置属性
var (
	GlobalAdminServerConfigProperties  = ServerConfigProperties{}
	GlobalPortalServerConfigProperties = ServerConfigProperties{}
	GlobalSearchServerConfigProperties = ServerConfigProperties{}
	GlobalJwtConfigProperties          = JwtConfigProperties{}
	// GlobalDatabaseConfigProperties 如果不同服务的数据库配置不同，可以在这里分别定义
	GlobalDatabaseConfigProperties = DatabaseConfigProperties{}
)

// ServerConfigProperties 服务配置属性
type ServerConfigProperties struct {
	// 启动
	Enable bool
	// release, debug, test
	GinRunMode      string
	Port            string
	ApplicationName string
	ReadTimeout     int
}

// JwtConfigProperties security 配置属性
type JwtConfigProperties struct {
	TokenHeader string
	TokenHead   string
	Secret      string
	Expiration  int
}

// GormMysqlConfigProperties gorm mysql 配置属性
type GormMysqlConfigProperties struct {
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool
	// 时区
	Loc string

	GormSlowThreshold             time.Duration
	GormColorful                  bool
	GormIgnoreRecordNotFoundError bool
	GormParameterizedQueries      bool
	// Gorm 日志级别
	// 1 Silent，2 Error，3 Warn，4 Info
	GormLogLevel int
}

// RedisConfigProperties redis 配置属性
type RedisConfigProperties struct {
	Host     string
	Port     string
	Password string
	Database int
}

type ElasticSearchConfigProperties struct {
	// host:port
	Addresses []string
	Username  string
	Password  string
	// http_ca证书路径 (http_ca.crt)
	CACertPath string
}

type DatabaseConfigProperties struct {
	GormMysqlConfigProperties     GormMysqlConfigProperties
	RedisConfigProperties         RedisConfigProperties
	ElasticSearchConfigProperties ElasticSearchConfigProperties
}
