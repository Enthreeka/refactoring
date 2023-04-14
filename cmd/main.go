package main

import (
	"github.com/Enthreeka/refactoring/internal/app/server"
	"github.com/Enthreeka/refactoring/pkg/logger"
)

func main() {

	log := logger.New()

	server.Run(log)
}
