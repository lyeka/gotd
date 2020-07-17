package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

const Bearer = "Bearer "

// Claims 本项目使用的 jwt 声明
type Claims struct {
	jwt.StandardClaims
	Uid int64 `json:"uid"` // 用户id
}

func (c *Claims) GetUserID() int64 {
	return c.Uid
}

// NewJwtToken 生成 jwt 令牌
// uid 用户id
// expiredAt 过期时间（Unix 时间戳）
func NewJwtToken(uid int64, expiredAt int64, key string) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiredAt},
		Uid:            uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))
	return Bearer + ss, err
}

// ParseToken
func ParseToken(tokenString string, key string) (*Claims, error) {
	// 去掉 Bearer
	s := strings.Fields(tokenString)
	if len(s) != 2 {
		return nil, errors.New("invalid token")
	}
	tokenString = s[1]
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalid")
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token assertion failed")
	}
}
