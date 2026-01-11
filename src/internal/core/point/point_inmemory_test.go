package point

import (
	"testing"

	"github.com/vishgupt/rainier/src/internal/models"
)

func TestUpsertPoints(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
	}

	err := manager.UpsertPoints("col1", points)
	if err != nil {
		t.Fatalf("UpsertPoints failed: %v", err)
	}
}

func TestGetPoints(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
	}

	manager.UpsertPoints("col1", points)

	retrieved, err := manager.GetPoints("col1", []string{"point1", "point2"})
	if err != nil {
		t.Fatalf("GetPoints failed: %v", err)
	}

	if len(retrieved) != 2 {
		t.Errorf("Expected 2 points, got %d", len(retrieved))
	}

	if retrieved[0].ID != "point1" {
		t.Errorf("Expected point ID 'point1', got '%s'", retrieved[0].ID)
	}
}

func TestGetPointsPartial(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
	}

	manager.UpsertPoints("col1", points)

	retrieved, err := manager.GetPoints("col1", []string{"point1"})
	if err != nil {
		t.Fatalf("GetPoints failed: %v", err)
	}

	if len(retrieved) != 1 {
		t.Errorf("Expected 1 point, got %d", len(retrieved))
	}

	if retrieved[0].ID != "point1" {
		t.Errorf("Expected point ID 'point1', got '%s'", retrieved[0].ID)
	}
}

func TestDeletePoints(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
	}

	manager.UpsertPoints("col1", points)

	err := manager.DeletePoints("col1", []string{"point1"})
	if err != nil {
		t.Fatalf("DeletePoints failed: %v", err)
	}

	retrieved, err := manager.GetPoints("col1", []string{"point1"})
	if err != nil {
		t.Fatalf("GetPoints failed: %v", err)
	}

	if len(retrieved) != 0 {
		t.Errorf("Expected 0 points after deletion, got %d", len(retrieved))
	}
}

func TestSearchPoints(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
		{
			ID:    "point3",
			Values: []float32{0.7, 0.8, 0.9},
		},
	}

	manager.UpsertPoints("col1", points)

	results, err := manager.SearchPoints("col1", []float32{0.1, 0.2, 0.3}, 2)
	if err != nil {
		t.Fatalf("SearchPoints failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestSearchPointsLimitGreaterThanTotal(t *testing.T) {
	manager := NewInMemoryManager()

	points := []*models.Point{
		{
			ID:    "point1",
			Values: []float32{0.1, 0.2, 0.3},
		},
		{
			ID:    "point2",
			Values: []float32{0.4, 0.5, 0.6},
		},
	}

	manager.UpsertPoints("col1", points)

	results, err := manager.SearchPoints("col1", []float32{0.1, 0.2, 0.3}, 10)
	if err != nil {
		t.Fatalf("SearchPoints failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}
