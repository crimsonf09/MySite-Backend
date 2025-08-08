package controller

import (
	"log"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/service"
	"github.com/gorilla/websocket"
)

func MessageHandler(conn *websocket.Conn, msg service.MessageInput) {
	msg.TimeStamp = time.Now()

	// IP Address not available directly through Gorilla WebSocket, leave blank or handle elsewhere
	msg.IPAddess = ""

	// Save to DB
	if _, err := service.GotNewMessage(msg); err != nil {
		log.Printf("Error saving message: %v", err)
		conn.WriteJSON(map[string]string{"error": "Failed to save message"})
		return
	}

	// Broadcast or process message
	processedMsg, err := service.GotNewMessage(msg)
	if err != nil {
		log.Printf("Error processing message: %v", err)
		conn.WriteJSON(map[string]string{"error": "Failed to process message"})
		return
	}

	// Send processed message back to client
	conn.WriteJSON(processedMsg)
}
