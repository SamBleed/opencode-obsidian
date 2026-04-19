package main

import (
	"github.com/SamBleed/ozy-api/internal/repository/db"
	"github.com/SamBleed/ozy-api/internal/server"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"os"
)

func main() {
	// Configuración (En producción usar ENV)
	bunkerSecret := "ozy-secret-2026"
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	// 1. Inicializar DB
	dbAdapter, err := db.NewPostgresAdapter(dbHost, "5432", "ozy_admin", "bunker_secret_2026", "ozy_db")
	if err != nil {
		logger.Error("Failed to connect to DB", "db-fail", map[string]interface{}{"error": err.Error()})
		// Permitimos que siga por ahora para no bloquear el dev si Docker está apagado
	} else {
		defer dbAdapter.Close()
	}

	// 2. Iniciar Servidor
	srv := server.NewServer(8080, bunkerSecret, dbAdapter)

	err = srv.Start()
	if err != nil {
		logger.Error("Failed to start server", "startup-fail", map[string]interface{}{"error": err.Error()})
		os.Exit(1)
	}
}
