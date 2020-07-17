package db

import (
	"context"
	"github.com/lyeka/gotd/internal/config"
	"github.com/lyeka/gotd/internal/mongo"
	_type "github.com/lyeka/gotd/internal/type"
)

type DB interface {
	// 创建用户
	CreateUser(ctx context.Context, user *_type.User) (id int64, err error)
	// 验证用户密码
	VerifyPassword(ctx context.Context, account, password string) (*_type.User, error) // todo account支持用户昵称，邮箱等登录

	// 创建任务
	CreateTask(ctx context.Context, task *_type.Task) (id int64, err error)
	// 更新任务
	UpdateTask(ctx context.Context, task *_type.Task) (err error)
	// 获取用户任务列表
	GetUserTasks(ctx context.Context, uid int64) ([]_type.Task, error)
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
