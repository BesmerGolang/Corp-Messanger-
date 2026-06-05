package main

import (
	"textMessanger/internal/auth"
	"textMessanger/internal/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := auth.NewRepository()
	authHandler := auth.NewHandler(repo)

	hub := ws.NewHub()
	go hub.Run()

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(hub, repo, c)
	})
	r.Run(":8080")
}
