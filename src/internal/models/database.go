package models

import (
	"time"

	"github.com/google/uuid"
)

// Database represents a database entity
type Database struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewDatabase creates a new database
func NewDatabase(name string) *Database {
	now := time.Now()
	return &Database{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
