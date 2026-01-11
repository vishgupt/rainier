package collection

import (
	"github.com/vishgupt/rainier/src/internal/models"
)

// Manager defines the interface for collection operations
type Manager interface {
	CreateCollection(databaseName, name string, dimension int32, metric string) (*models.Collection, error)
	GetCollection(databaseName, name string) (*models.Collection, error)
	ListCollections(databaseName string, page, limit int) ([]*models.Collection, int64, error)
	DeleteCollection(databaseName, name string) error
}
