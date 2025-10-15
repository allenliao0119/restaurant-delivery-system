package main

import (
	"log"

	"github.com/allenliao0119/restaurant-delivery-system/internal/server"
	"github.com/allenliao0119/restaurant-delivery-system/pkg/config"
	"github.com/allenliao0119/restaurant-delivery-system/pkg/db"
)

func main() {
	cfg := config.Load()
	engine, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	s := server.New(engine)
	log.Printf("🚀 Starting API server on :%s", cfg.Port)
	if err := s.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server start failed: %v", err)
	}
}
