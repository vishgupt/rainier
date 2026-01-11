package collection

import (
	"testing"
)

func TestCreateCollection(t *testing.T) {
	manager := NewInMemoryManager()

	col, err := manager.CreateCollection("test_db", "test_col", 128, "euclidean")
	if err != nil {
		t.Fatalf("CreateCollection failed: %v", err)
	}

	if col.Name != "test_col" {
		t.Errorf("Expected collection name 'test_col', got '%s'", col.Name)
	}

	if col.DatabaseName != "test_db" {
		t.Errorf("Expected database name 'test_db', got '%s'", col.DatabaseName)
	}

	if col.Dimension != 128 {
		t.Errorf("Expected dimension 128, got %d", col.Dimension)
	}

	if col.Metric != "euclidean" {
		t.Errorf("Expected metric 'euclidean', got '%s'", col.Metric)
	}
}

func TestGetCollection(t *testing.T) {
	manager := NewInMemoryManager()

	created, _ := manager.CreateCollection("test_db", "test_col", 128, "euclidean")

	retrieved, err := manager.GetCollection("test_db", "test_col")
	if err != nil {
		t.Fatalf("GetCollection failed: %v", err)
	}

	if retrieved.ID != created.ID {
		t.Errorf("Expected collection ID %s, got %s", created.ID, retrieved.ID)
	}

	if retrieved.Name != "test_col" {
		t.Errorf("Expected collection name 'test_col', got '%s'", retrieved.Name)
	}
}

func TestListCollections(t *testing.T) {
	manager := NewInMemoryManager()

	manager.CreateCollection("test_db", "col1", 128, "euclidean")
	manager.CreateCollection("test_db", "col2", 256, "cosine")
	manager.CreateCollection("test_db", "col3", 512, "dot_product")

	collections, total, err := manager.ListCollections("test_db", 1, 10)
	if err != nil {
		t.Fatalf("ListCollections failed: %v", err)
	}

	if total != 3 {
		t.Errorf("Expected 3 collections, got %d", total)
	}

	if len(collections) != 3 {
		t.Errorf("Expected 3 collections in result, got %d", len(collections))
	}
}

func TestListCollectionsPagination(t *testing.T) {
	manager := NewInMemoryManager()

	for i := 1; i <= 5; i++ {
		manager.CreateCollection("test_db", "col"+string(rune(48+i)), 128, "euclidean")
	}

	collections, total, err := manager.ListCollections("test_db", 1, 2)
	if err != nil {
		t.Fatalf("ListCollections failed: %v", err)
	}

	if total != 5 {
		t.Errorf("Expected total 5, got %d", total)
	}

	if len(collections) != 2 {
		t.Errorf("Expected 2 collections in page 1, got %d", len(collections))
	}
}

func TestDeleteCollection(t *testing.T) {
	manager := NewInMemoryManager()

	manager.CreateCollection("test_db", "test_col", 128, "euclidean")

	err := manager.DeleteCollection("test_db", "test_col")
	if err != nil {
		t.Fatalf("DeleteCollection failed: %v", err)
	}

	_, err = manager.GetCollection("test_db", "test_col")
	if err == nil {
		t.Error("Expected GetCollection to fail after deletion")
	}
}
