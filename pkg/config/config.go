package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Config loaded. Port: %s\n", port)
	return Config{Port: port}
}
