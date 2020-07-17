package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gotd/internal/config"
	"net/http"
)

const UID = "uid"

// Auth 认证中间件
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/api/v1/user/login" || c.Request.RequestURI == "/api/v1/user/register" {
			c.Next()
			return
		}

		tokenString := c.Request.Header.Get("Authorization")
		// 登录才可以使用
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "请登录后使用"})
			c.Abort()
		}
		claims, err := ParseToken(tokenString, cfg.JWTKey)
		if err != nil {
			fmt.Println("parse jwt token failed, error: ", err)
			c.Abort()
		}
		uid := claims.GetUserID()
		c.Set(UID, uid)

		c.Next()

		// todo 延长 token 有效期
	}
}
