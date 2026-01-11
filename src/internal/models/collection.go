package models

import (
	"time"

	"github.com/google/uuid"
)

// Collection represents a vector collection
type Collection struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	DatabaseName string            `json:"database_name"`
	Dimension    int32             `json:"dimension"`
	Metric       string            `json:"metric"`
	IndexConfig  map[string]interface{} `json:"index_config,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

// NewCollection creates a new collection
func NewCollection(name, databaseName string, dimension int32, metric string) *Collection {
	now := time.Now()
	return &Collection{
		ID:           uuid.New().String(),
		Name:         name,
		DatabaseName: databaseName,
		Dimension:    dimension,
		Metric:       metric,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
