package database

import (
	"github.com/vishgupt/rainier/src/internal/models"
)

// Manager defines the interface for database operations
type Manager interface {
	CreateDatabase(name string) (*models.Database, error)
	GetDatabase(name string) (*models.Database, error)
	ListDatabases(page, limit int) ([]*models.Database, int64, error)
	DeleteDatabase(name string) error
}
