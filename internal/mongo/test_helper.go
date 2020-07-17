package mongo

import (
	"context"
	"github.com/lyeka/gotd/internal/config"
)

func getTestDB() *DB {
	ctx := context.Background()
	cfg := new(config.Config)
	cfg.MongoDB = config.MongoDB{
		User:     "myUserAdmin",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     "27017",
		DB:       "gotd",
		AuthDB:   "admin",
	}
	testDB, _ := OpenDB(ctx, cfg)
	return testDB
}
