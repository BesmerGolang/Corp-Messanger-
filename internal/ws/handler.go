package ws

import (
	"net/http"

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
