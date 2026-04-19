package domain

import (
	"time"
)

// Entity representa un agente o herramienta registrada en el Bunker
type Entity struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Status    string    `json:"status" db:"status"`
	Metadata  string    `json:"metadata" db:"metadata"` // JSON string
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// EntityRepository define el puerto para la persistencia
type EntityRepository interface {
	Save(e *Entity) error
	List() ([]Entity, error)
	GetByName(name string) (*Entity, error)
}
