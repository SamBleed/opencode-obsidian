package db

import (
	"fmt"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	DB *sqlx.DB
}

func NewPostgresAdapter(host, port, user, password, dbname string) (*PostgresAdapter, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	logger.Info("Connected to PostgreSQL Bunker Database", map[string]interface{}{
		"host": host,
		"port": port,
		"db":   dbname,
	})

	return &PostgresAdapter{DB: db}, nil
}

func (a *PostgresAdapter) Close() error {
	return a.DB.Close()
}
