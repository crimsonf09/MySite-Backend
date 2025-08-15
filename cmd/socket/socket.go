package socket

import (
	"log"
	"net/http"

	"github.com/crimsonf09/MySite-Backend/internal/controller"
	"github.com/crimsonf09/MySite-Backend/internal/service"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for development, restrict in production
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	for {
		var msg service.MessageInput
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		// Handle message (already handles sending response)
		controller.MessageHandler(conn, msg)

	}
}
