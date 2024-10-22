package main

import (
	"log"

	"github.com/souvik03-136/Fam-Go/internal/config"
	"github.com/souvik03-136/Fam-Go/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
		return
	}

	s := server.NewServer()

	if err := s.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer cfg.CloseDB()
}
