package main

import (
	"log"

	"github.com/crimsonf09/MySite-Backend/cmd/api"
)

func main() {
	if err := api.StartAPIServer(); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
