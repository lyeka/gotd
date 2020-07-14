package server

import (
	"github.com/gin-gonic/gin"
)

func Ping(s *Server, group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": s.EX,
		})
	})
}
