package point

import (
	"sync"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/models"
)

// InMemoryManager implements Manager interface with in-memory storage
type InMemoryManager struct {
	mu     sync.RWMutex
	points map[string]map[string]*models.Point
}

// NewInMemoryManager creates a new in-memory point manager
func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		points: make(map[string]map[string]*models.Point),
	}
}

// UpsertPoints inserts or updates points
func (m *InMemoryManager) UpsertPoints(collectionID string, points []*models.Point) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.points[collectionID]; !exists {
		m.points[collectionID] = make(map[string]*models.Point)
	}

	for _, point := range points {
		m.points[collectionID][point.ID] = point
	}

	return nil
}

// GetPoints retrieves points by IDs
func (m *InMemoryManager) GetPoints(collectionID string, ids []string) ([]*models.Point, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	collection, exists := m.points[collectionID]
	if !exists {
		return nil, common.NewNotFoundError("collection not found")
	}

	result := make([]*models.Point, 0, len(ids))
	for _, id := range ids {
		if point, exists := collection[id]; exists {
			result = append(result, point)
		}
	}

	return result, nil
}

// DeletePoints deletes points by IDs
func (m *InMemoryManager) DeletePoints(collectionID string, ids []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	collection, exists := m.points[collectionID]
	if !exists {
		return common.NewNotFoundError("collection not found")
	}

	for _, id := range ids {
		delete(collection, id)
	}

	return nil
}

// SearchPoints performs a simple search (placeholder for vector similarity)
func (m *InMemoryManager) SearchPoints(collectionID string, vector []float32, topK int) ([]*models.Point, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	collection, exists := m.points[collectionID]
	if !exists {
		return nil, common.NewNotFoundError("collection not found")
	}

	result := make([]*models.Point, 0, topK)
	count := 0
	for _, point := range collection {
		if count >= topK {
			break
		}
		result = append(result, point)
		count++
	}

	return result, nil
}
