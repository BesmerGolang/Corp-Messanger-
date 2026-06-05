package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	userID int
}

func (c *Client) readPump() {
	defer func() { //а эта штука для того, что бы сервер не пытался отправлять закрытому клиенту сообщения после ошибки в цикле
		c.hub.unregister <- c //отправили челика в канал на удаление
		c.conn.Close()        //закрыли вебсокет соединение
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка чтения сообщений из контекста webSocket")
			break
		}
		c.hub.broadcast <- message //запомнить, вот это просто отправка в канал
	}
}
func (c *Client) writePump() {
	for {
		msg, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			break
		}
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Ошибка отправки сообщения клиенту")
			break
		}
	}
}
