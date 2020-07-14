package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// Config 配置
type Config struct {
	Port string // 运行端口
	DBEngine string // 数据库源

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
func (cfg *Config) DbDSN() string {
	switch cfg.DBEngine {
	case "MongoDB":
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.DB, cfg.AuthDB)
	case "Redis":
		return ""
	case "MySQL":
		return ""
	default: // 默认mongodb
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.DB, cfg.AuthDB)
	}
}
