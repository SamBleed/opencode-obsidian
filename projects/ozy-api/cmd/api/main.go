package main

import (
	"github.com/SamBleed/ozy-api/internal/server"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"os"
)

func main() {
	srv := server.NewServer(8080)

	err := srv.Start()
	if err != nil {
		logger.Error("Failed to start server", "startup-fail", map[string]interface{}{"error": err.Error()})
		os.Exit(1)
	}
}
