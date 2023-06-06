package conf

import "time"

// 全局配置属性
var (
	GlobalServerConfigProperties   = ServerConfigProperties{}
	GlobalJwtConfigProperties      = JwtConfigProperties{}
	GlobalDatabaseConfigProperties = DatabaseConfigProperties{}
)

// ServerConfigProperties 服务配置属性
type ServerConfigProperties struct {
	GinRunMode  string
	Host        string
	Port        string
	ReadTimeout int64
}

// JwtConfigProperties jwt 配置属性
type JwtConfigProperties struct {
	TokenHeader string
	TokenHead   string
	Secret      string
	Expiration  int64
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

type RedisConfigProperties struct {
	Host     string
	Port     string
	Password string
	Database int
}

type DatabaseConfigProperties struct {
	GormMysqlConfigProperties GormMysqlConfigProperties
	RedisConfigProperties     RedisConfigProperties
}
