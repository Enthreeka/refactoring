package main

import (
	"github.com/Enthreeka/refactoring/internal/app/server"
	"github.com/Enthreeka/refactoring/internal/config"
	"github.com/Enthreeka/refactoring/pkg/logger"
)

func main() {

	log := logger.New()

	configPath := "config.json"

	config, err := config.New(configPath)
	if err != nil {
		log.Fatal("Failed to load config: %s", err)
	}
	server.Run(config, log)
}
