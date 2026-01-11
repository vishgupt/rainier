package database

import (
	"sync"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/models"
)

// InMemoryManager implements Manager interface with in-memory storage
type InMemoryManager struct {
	mu        sync.RWMutex
	databases map[string]*models.Database
}

// NewInMemoryManager creates a new in-memory database manager
func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		databases: make(map[string]*models.Database),
	}
}

// CreateDatabase creates a new database
func (m *InMemoryManager) CreateDatabase(name string) (*models.Database, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.databases[name]; exists {
		return nil, common.NewValidationError("database already exists")
	}

	db := models.NewDatabase(name)
	m.databases[name] = db
	return db, nil
}

// GetDatabase retrieves a database by name
func (m *InMemoryManager) GetDatabase(name string) (*models.Database, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	db, exists := m.databases[name]
	if !exists {
		return nil, common.NewNotFoundError("database not found")
	}
	return db, nil
}

// ListDatabases lists all databases with pagination
func (m *InMemoryManager) ListDatabases(page, limit int) ([]*models.Database, int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	total := int64(len(m.databases))
	start := (page - 1) * limit
	end := start + limit

	if start >= int(total) {
		return []*models.Database{}, total, nil
	}

	if end > int(total) {
		end = int(total)
	}

	result := make([]*models.Database, 0, limit)
	i := 0
	for _, db := range m.databases {
		if i >= start && i < end {
			result = append(result, db)
		}
		i++
	}

	return result, total, nil
}

// DeleteDatabase deletes a database
func (m *InMemoryManager) DeleteDatabase(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.databases[name]; !exists {
		return common.NewNotFoundError("database not found")
	}

	delete(m.databases, name)
	return nil
}
