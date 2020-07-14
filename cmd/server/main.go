package main

import server2 "github.com/lyeka/gotd/internal/server"

func main() {
	server := server2.NewServer()
	server.Run()
}