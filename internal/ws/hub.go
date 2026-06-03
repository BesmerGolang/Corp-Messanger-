package ws

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					delete(h.clients, client)
					close(client.send)
				}
			}
		}
	}
}

/*Попробуй написать его. Логика там будет такая:

Поймать клиента из канала h.unregister (точно так же, как ты сделал с register).
Проверить, есть ли он в мапе (опционально, но желательно для безопасности).
Удалить его из мапы с помощью встроенной функции delete(мапа, ключ).
Закрыть личный канал клиента, чтобы освободить ресурсы (функция close()). Предположим, что в структуре Client его канал называется send.*/
