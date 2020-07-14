package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gotd/internal/config"
	"github.com/lyeka/gotd/internal/db"
	"log"
)

// Server ...
type Server struct {
	DB db.DB
	Engine *gin.Engine
	EX string // 测试字段
}

func (s *Server) Run(port string) {
	log.Fatal(s.Engine.Run(port))
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

func NewServer(cfg *config.Config) *Server {
	server := &Server{}
	server.Engine = newEngine()

	ctx := context.Background()
	ddb, err := db.OpenDB(ctx, cfg)
	if err != nil {
		log.Fatal("connect db failed", err)
	}
	server.DB = ddb

	server.EX = "go todo"

	registerRouter(server)

	return server
}




