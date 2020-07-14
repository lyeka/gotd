package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gotd/internal/db"
	"log"
)

// Server ...
type Server struct {
	DB db.DB
	Engine *gin.Engine
	EX string // 测试字段
}

func (s *Server) Run() {
	log.Fatal(s.Engine.Run(":8080"))
}

func newEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	return r
}

// 注册路由
func registerRouter(s *Server) {

	v1 := s.Engine.Group("/api/v1")
	{
		Ping(s, v1)
	}
}

func NewServer() *Server {
	server := &Server{}
	server.Engine = newEngine()
	server.EX = "go todo"

	registerRouter(server)

	return server
}




