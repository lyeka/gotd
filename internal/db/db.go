package db

import (
	"context"
	"github.com/lyeka/gotd/internal/config"
	"github.com/lyeka/gotd/internal/mongo"
	_type "github.com/lyeka/gotd/internal/type"
)

type DB interface {
	CreateUser(ctx context.Context, user *_type.User) (id string,  err error)
	VerifyPassword(ctx context.Context, account, password string) (*_type.User, error)// todo account支持用户昵称，邮箱等登录
}

func OpenDB(ctx context.Context, cfg *config.Config) (db DB, err error) {
	switch cfg.DBEngine {
	case "MongoDB":
		db, err = mongo.OpenDB(ctx, cfg)
	default:
		db, err = mongo.OpenDB(ctx, cfg)
	}
	return
}