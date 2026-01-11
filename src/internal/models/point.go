package models

import (
	"time"

	"github.com/google/uuid"
)

// Point represents a vector point
type Point struct {
	ID           string                 `json:"id"`
	CollectionID string                 `json:"collection_id"`
	Values       []float32              `json:"values"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

// NewPoint creates a new point
func NewPoint(collectionID string, values []float32, metadata map[string]interface{}) *Point {
	now := time.Now()
	return &Point{
		ID:           uuid.New().String(),
		CollectionID: collectionID,
		Values:       values,
		Metadata:     metadata,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
