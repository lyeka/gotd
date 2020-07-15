package mongo

import (
	"context"
	"errors"
	_type "github.com/lyeka/gotd/internal/type"
	"github.com/lyeka/gotd/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Email string
	Nickname string
	Password string
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}


// Register 创建用户
func (db *DB) CreateUser(ctx context.Context, user *_type.User) (id string, err error) {
	exist, err := db.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return
	}
	if exist != nil {
		err = errors.New("用户邮箱已存在")
		return
	}

	now := time.Now()
	// 密码加密
	pwd, err := pkg.HashAndSalt(user.Password)
	if err != nil {
		return
	}
	innerUser := &User{
		ID:        primitive.NewObjectID(), // todo NewObjectID 方法是基于当前时间生成的 id，存在不唯一问题
		Email:     user.Email,
		Nickname:  user.Nickname,
		Password:  pwd,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := db.CollUser().InsertOne(ctx, innerUser)
	if err != nil {
		return
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()

	return
}

// VerifyPassword 验证密码
// 正确返回用户信息
func (db *DB) VerifyPassword(ctx context.Context, email, password string) (*_type.User, error) {
	user, err := db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !pkg.ComparePassword(user.Password, password) {
		return nil, errors.New("password incorrect")
	}

	return &_type.User{
		Id: user.ID.Hex(),
		Nickname: user.Nickname,
		Email: user.Email,
	}, nil
}

// GetUserByEmail 通过邮箱查询用户
func (db *DB) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	result := db.CollUser().FindOne(ctx, bson.M{"email": email})
	if err := result.Err(); err != nil {
		return nil, err
	}

	user := new(User)
	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
