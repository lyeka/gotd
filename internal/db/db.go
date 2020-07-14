package db

import _type "github.com/lyeka/gotd/internal/type"

type DB interface {
	Register(user _type.User) error
	Login(account, password string) // account支持用户昵称，邮箱等登录
}