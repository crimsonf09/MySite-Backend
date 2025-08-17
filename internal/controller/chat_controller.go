package controller

import (
	"log"
	"time"

	"github.com/crimsonf09/MySite-Backend/internal/service"
	"github.com/gorilla/websocket"
)

func MessageHandler(conn *websocket.Conn, msg service.MessageInput) {
	// Add timestamp & IP
	msg.TimeStamp = time.Now()
	msg.IPAddress = "" // optional: extract from request

	// Save user message & get bot response
	_, botMsg, err := service.GotNewMessage(msg)
	if err != nil {
		log.Printf("Error saving message: %v", err)
		conn.WriteJSON(map[string]string{"error": "Failed to save message"})
		return
	}

	// Send bot message back to client
	if err := conn.WriteJSON(map[string]string{
		"message": botMsg.Message,
		"uid":     botMsg.UID,
		"sender":  botMsg.Sender,
	}); err != nil {
		log.Printf("Error sending bot message: %v", err)
		return
	}
}
