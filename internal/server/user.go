package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gotd/internal/auth"
	"github.com/lyeka/gotd/internal/config"
	_type "github.com/lyeka/gotd/internal/type"
	"net/http"
	"time"
)

// Register 注册用户
func Register(s *Server, rg *gin.RouterGroup) {
	rg.POST("/register", func(c *gin.Context) {
		resp := &Response{Code: http.StatusOK}
		nickName := c.PostForm("nickname")
		email := c.PostForm("email")
		password := c.PostForm("password")

		user := &_type.User{
			Nickname: nickName,
			Email: email,
			Password: password,
		}

		uid, err := s.DB.CreateUser(c, user)
		if err != nil {
			// todo code
			resp.Message = err.Error()
			c.JSON(200, resp)
			return
		}

		resp.Message = "注册成功"
		resp.Data = uid
		c.JSON(200, resp)
		return
	})
}

// Login 用户登录
func Login(s *Server, cfg *config.Config, rg *gin.RouterGroup)  {
	rg.POST("/login", func(c *gin.Context) {
		resp := &Response{Code: http.StatusOK}
		email := c.PostForm("email")
		pwd := c.PostForm("password")

		user, err := s.DB.VerifyPassword(c, email, pwd)
		if err != nil {
			// todo code
			resp.Message = "密码错误"
			c.JSON(http.StatusOK, resp)
			return
		}

		uid := user.Id
		expiredAt := time.Now().Unix() + cfg.JWTTtl
		token, err := auth.NewJwtToken(uid, expiredAt, cfg.JWTKey)
		if err != nil {
			// todo code
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = token
		c.JSON(http.StatusOK, resp)
		return
	})
}

