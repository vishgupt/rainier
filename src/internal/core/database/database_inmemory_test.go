package database

import (
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	manager := NewInMemoryManager()

	db, err := manager.CreateDatabase("test_db")
	if err != nil {
		t.Fatalf("CreateDatabase failed: %v", err)
	}

	if db.Name != "test_db" {
		t.Errorf("Expected database name 'test_db', got '%s'", db.Name)
	}

	if db.ID == "" {
		t.Error("Expected database ID to be set")
	}
}

func TestGetDatabase(t *testing.T) {
	manager := NewInMemoryManager()

	created, _ := manager.CreateDatabase("test_db")

	retrieved, err := manager.GetDatabase("test_db")
	if err != nil {
		t.Fatalf("GetDatabase failed: %v", err)
	}

	if retrieved.ID != created.ID {
		t.Errorf("Expected database ID %s, got %s", created.ID, retrieved.ID)
	}

	if retrieved.Name != "test_db" {
		t.Errorf("Expected database name 'test_db', got '%s'", retrieved.Name)
	}
}

func TestListDatabases(t *testing.T) {
	manager := NewInMemoryManager()

	manager.CreateDatabase("db1")
	manager.CreateDatabase("db2")
	manager.CreateDatabase("db3")

	databases, total, err := manager.ListDatabases(1, 10)
	if err != nil {
		t.Fatalf("ListDatabases failed: %v", err)
	}

	if total != 3 {
		t.Errorf("Expected 3 databases, got %d", total)
	}

	if len(databases) != 3 {
		t.Errorf("Expected 3 databases in result, got %d", len(databases))
	}
}

func TestListDatabasesPagination(t *testing.T) {
	manager := NewInMemoryManager()

	for i := 1; i <= 5; i++ {
		manager.CreateDatabase("db" + string(rune(48+i)))
	}

	databases, total, err := manager.ListDatabases(1, 2)
	if err != nil {
		t.Fatalf("ListDatabases failed: %v", err)
	}

	if total != 5 {
		t.Errorf("Expected total 5, got %d", total)
	}

	if len(databases) != 2 {
		t.Errorf("Expected 2 databases in page 1, got %d", len(databases))
	}
}

func TestDeleteDatabase(t *testing.T) {
	manager := NewInMemoryManager()

	manager.CreateDatabase("test_db")

	err := manager.DeleteDatabase("test_db")
	if err != nil {
		t.Fatalf("DeleteDatabase failed: %v", err)
	}

	_, err = manager.GetDatabase("test_db")
	if err == nil {
		t.Error("Expected GetDatabase to fail after deletion")
	}
}
