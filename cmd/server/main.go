package main

import (
	"textMessanger/internal/auth"
	"textMessanger/internal/database"
	"textMessanger/internal/ws"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) { //кароче пока сам не до конца понял, но это разрешения для CORS, дабы он не блочил нам пост запросы между сервером и сайтом
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" { //Вот такой запрос он отправляет что бы проверить, все ли гуд с сервером, поэтому на него автоматом отвечает 204 типо все гуд
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	db := database.InitDB()
	defer db.Close()

	repo := auth.NewRepository(db)
	authHandler := auth.NewHandler(repo)

	hub := ws.NewHub()
	go hub.Run()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(CORSMiddleware())
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.GET("/ws", auth.AuthMiddleware(), func(c *gin.Context) {
		ws.ServeWs(hub, c)
	})
	r.Run(":8080")
}
