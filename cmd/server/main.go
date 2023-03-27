package main

import (
	"log"

	"github.com/hmdsefi/word-of-wisdom/config"
	"github.com/hmdsefi/word-of-wisdom/internal/server/app"
)

func main() {
	// Configuration
	cfg, err := config.NewServerConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
