package server

import (
	"github.com/gin-gonic/gin"
	_type "github.com/lyeka/gotd/internal/type"
)

// Register 注册用户
func Register(s *Server, g *gin.RouterGroup) {
	g.POST("/register", func(c *gin.Context) {
		nickName := c.PostForm("nickName")
		email := c.PostForm("email")
		password := c.PostForm("password")

		user := &_type.User{
			NickName: nickName,
			Email: email,
			Password: password,
		}

		err := s.DB.Register(user)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "注册失败",
			})
		}

	})
}

// Login 用户登录
func Login(s *Server)  {

}

