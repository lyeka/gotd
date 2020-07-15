package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// Config 配置
type Config struct {
	Port string // 运行端口
	DBEngine string // 数据库源
	JWTKey string // jwt 签名
	JWTTtl int64 // jwt token 有效期时长单位s
	MongoDB
}

type MongoDB struct {
	User string
	Password string
	Host string
	Port string
	DB string
	AuthDB string
}

// Init 初始化配置
func Init(cfgPath string) (*Config, error) {
	cfg := new(Config)
	err := ini.MapTo(cfg,cfgPath)
	return cfg, err
}


func (cfg *Config) RunPort() string {
	return ":" + cfg.Port
}

// DnDSN 返回数据库DSN
func (cfg *Config) DbDSN() (db, dsn string) {
	switch cfg.DBEngine {
	case "MongoDB":
		db =  cfg.MongoDB.DB
		dsn = fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
			cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.AuthDB)
	case "Redis":

	case "MySQL":
	default: // 默认mongodb
		db =  cfg.MongoDB.DB
		dsn = fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
			cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.AuthDB)
	}
	return
}
