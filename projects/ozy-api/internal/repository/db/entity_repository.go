package db

import (
	"github.com/SamBleed/ozy-api/internal/domain"
)

type SqlEntityRepository struct {
	adapter *PostgresAdapter
}

func NewSqlEntityRepository(adapter *PostgresAdapter) *SqlEntityRepository {
	return &SqlEntityRepository{adapter: adapter}
}

func (r *SqlEntityRepository) Save(e *domain.Entity) error {
	query := `
		INSERT INTO entities (name, status, metadata, updated_at)
		VALUES (:name, :status, :metadata, NOW())
		ON CONFLICT (name) DO UPDATE 
		SET status = EXCLUDED.status, metadata = EXCLUDED.metadata, updated_at = NOW()
	`
	_, err := r.adapter.DB.NamedExec(query, e)
	return err
}

func (r *SqlEntityRepository) List() ([]domain.Entity, error) {
	var entities []domain.Entity
	query := `SELECT id, name, status, metadata, updated_at FROM entities ORDER BY updated_at DESC`
	err := r.adapter.DB.Select(&entities, query)
	return entities, err
}

func (r *SqlEntityRepository) GetByName(name string) (*domain.Entity, error) {
	var e domain.Entity
	query := `SELECT id, name, status, metadata, updated_at FROM entities WHERE name = $1`
	err := r.adapter.DB.Get(&e, query, name)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
