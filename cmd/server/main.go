package main

import (
	"textMessanger/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authRepo := auth.NewRepisitory()
	authHandler := auth.NewHandler(authRepo)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
	r.Run(":8080")
}
