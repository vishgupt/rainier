package collection

import (
	"sync"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/models"
)

// InMemoryManager implements Manager interface with in-memory storage
type InMemoryManager struct {
	mu          sync.RWMutex
	collections map[string]map[string]*models.Collection
}

// NewInMemoryManager creates a new in-memory collection manager
func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		collections: make(map[string]map[string]*models.Collection),
	}
}

// CreateCollection creates a new collection
func (m *InMemoryManager) CreateCollection(databaseName, name string, dimension int32, metric string) (*models.Collection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.collections[databaseName]; !exists {
		m.collections[databaseName] = make(map[string]*models.Collection)
	}

	if _, exists := m.collections[databaseName][name]; exists {
		return nil, common.NewValidationError("collection already exists")
	}

	collection := models.NewCollection(name, databaseName, dimension, metric)
	m.collections[databaseName][name] = collection
	return collection, nil
}

// GetCollection retrieves a collection by name
func (m *InMemoryManager) GetCollection(databaseName, name string) (*models.Collection, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	db, exists := m.collections[databaseName]
	if !exists {
		return nil, common.NewNotFoundError("database not found")
	}

	collection, exists := db[name]
	if !exists {
		return nil, common.NewNotFoundError("collection not found")
	}

	return collection, nil
}

// ListCollections lists all collections in a database with pagination
func (m *InMemoryManager) ListCollections(databaseName string, page, limit int) ([]*models.Collection, int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	db, exists := m.collections[databaseName]
	if !exists {
		return []*models.Collection{}, 0, nil
	}

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	total := int64(len(db))
	start := (page - 1) * limit
	end := start + limit

	if start >= int(total) {
		return []*models.Collection{}, total, nil
	}

	if end > int(total) {
		end = int(total)
	}

	result := make([]*models.Collection, 0, limit)
	i := 0
	for _, collection := range db {
		if i >= start && i < end {
			result = append(result, collection)
		}
		i++
	}

	return result, total, nil
}

// DeleteCollection deletes a collection
func (m *InMemoryManager) DeleteCollection(databaseName, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	db, exists := m.collections[databaseName]
	if !exists {
		return common.NewNotFoundError("database not found")
	}

	if _, exists := db[name]; !exists {
		return common.NewNotFoundError("collection not found")
	}

	delete(db, name)
	return nil
}
