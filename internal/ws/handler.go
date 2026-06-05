package ws

import (
	"fmt"
	"net/http"
	"textMessanger/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// тут будет web socket(надеюсь)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024, //это настройки буфера для чтения и отправки файла, стандартный размер 1кб, если выставить слишком маленький, а сообщения будут большими - потеряем производительность
	CheckOrigin: func(r *http.Request) bool { //тут мы разрешаем подключение любых доменов
		return true
	},
}

func ServeWs(hub *Hub, repo *auth.Repository, c *gin.Context) {
	userName := c.Query("username") // получаем имя пользователя из query параметров
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found"})
		return
	}

	user, err := repo.GetUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	send := make(chan []byte, 256)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("ошибка при апгрейде соединения", err)
		return
	}

	newClient := &Client{
		hub:    hub,
		conn:   conn,
		send:   send,
		userID: user.ID,
	}

	hub.register <- newClient
	go newClient.writePump()
	go newClient.readPump()
}
