package main

import (
	"log"
	"github.com/allenliao0119/restaurant-delivery-system/internal/server"
)

func main() {
	s := server.New()
	log.Println("🚀 Starting API server on :8080")
	if err := s.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
