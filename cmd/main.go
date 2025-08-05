package main

import (
	"log"

	"github.com/crimsonf09/MySite-Backend/cmd/api"
	"github.com/crimsonf09/MySite-Backend/cmd/socket"
)

func main() {
	go socket.StartSocketServer() // WebSocket or TCP server

	if err := api.StartAPIServer(); err != nil {
		log.Fatalf("API server failed: %v", err)
	}
}
