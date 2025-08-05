package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins, use stricter checks in prod
	},
}

func StartSocketServer() {
	http.HandleFunc("/ws", handleSocket)

	log.Println("Socket server listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Socket server failed: %v", err)
	}
}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Socket read error:", err)
			break
		}
		log.Println("Received from socket:", string(msg))
		conn.WriteMessage(websocket.TextMessage, []byte("pong"))
	}
}
