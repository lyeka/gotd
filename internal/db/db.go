package db

import (
	"context"
	"github.com/lyeka/gotd/internal/config"
	"github.com/lyeka/gotd/internal/mongo"
	_type "github.com/lyeka/gotd/internal/type"
)

type DB interface {
	Register(user *_type.User) error
	Login(account, password string) error// account支持用户昵称，邮箱等登录
}

func OpenDB(ctx context.Context, cfg *config.Config) (db DB, err error) {
	switch cfg.DBEngine {
	case "MongoDB":
		db, err = mongo.OpenDB(ctx, cfg.DbDSN())
	default:

	}
	return
}