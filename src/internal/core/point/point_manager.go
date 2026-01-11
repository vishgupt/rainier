package point

import (
	"github.com/vishgupt/rainier/src/internal/models"
)

// Manager defines the interface for point operations
type Manager interface {
	UpsertPoints(collectionID string, points []*models.Point) error
	GetPoints(collectionID string, ids []string) ([]*models.Point, error)
	DeletePoints(collectionID string, ids []string) error
	SearchPoints(collectionID string, vector []float32, topK int) ([]*models.Point, error)
}
