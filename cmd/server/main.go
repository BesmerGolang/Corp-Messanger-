package main

import (
	"net/http"
	"textMessanger/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authRepo := auth.NewRepository()
	authHandler := auth.NewHandler(authRepo)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
	// Приватные маршруты (требуют токен)
	protected := r.Group("/api")
	protected.Use(auth.AuthMiddleware())
	{
		// Пример: получить информацию о текущем пользователе
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			// Здесь можно сходить в репозиторий и вернуть данные пользователя
			c.JSON(http.StatusOK, gin.H{"user_id": userID})
		})
	}
	r.Run(":8080")
}
