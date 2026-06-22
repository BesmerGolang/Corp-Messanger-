package ws

import (
	"fmt"
	"net/http"

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

func ServeWs(hub *Hub, c *gin.Context) {
	userIDAny, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "пользователь не авторизован"})
		return
	}
	userID := userIDAny.(int)
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
		userID: userID,
	}
	history, err := hub.repo.GetLastMessages()
	if err != nil {
		for _, msg := range history {
			newClient.conn.WriteJSON(msg)
		}
	} else {
		fmt.Println("Ошибка загрузки истории")
	}

	hub.register <- newClient
	go newClient.writePump()
	go newClient.readPump()
}
