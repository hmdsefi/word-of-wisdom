package main

import (
	"log"

	"github.com/hmdsefi/word-of-wisdom/config"
	"github.com/hmdsefi/word-of-wisdom/internal/client/app"
)

func main() {
	// Configuration
	cfg, err := config.NewClientConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
