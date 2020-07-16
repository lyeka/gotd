package main

import (
	"github.com/lyeka/gotd/internal/config"
	server2 "github.com/lyeka/gotd/internal/server"
	"log"
)

func main() {
	cfg, err := config.Init("./config.ini")
	if err != nil {
		log.Fatal("init config failed")
		return
	}

	server := server2.NewServer(cfg)
	server.Run(cfg.RunPort())
}
