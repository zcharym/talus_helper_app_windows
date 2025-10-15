package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/models"
)

// Storage interface defines methods for data persistence
type Storage interface {
	LoadTodos() ([]models.Todo, error)
	SaveTodos(todos []models.Todo) error
}

// JSONStorage implements Storage interface using JSON files
type JSONStorage struct {
	dataDir string
}

// NewJSONStorage creates a new JSON storage instance
func NewJSONStorage() *JSONStorage {
	return &JSONStorage{}
}

// LoadTodos loads todos from the JSON file
func (s *JSONStorage) LoadTodos() ([]models.Todo, error) {
	dataDir, err := config.GetDataDir()
	if err != nil {
		return nil, err
	}

	todosFile := filepath.Join(dataDir, "todos.json")
	data, err := os.ReadFile(todosFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Todo{}, nil
		}
		return nil, err
	}

	var todos []models.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// SaveTodos saves todos to the JSON file
func (s *JSONStorage) SaveTodos(todos []models.Todo) error {
	dataDir, err := config.GetDataDir()
	if err != nil {
		return err
	}

	todosFile := filepath.Join(dataDir, "todos.json")
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(todosFile, data, 0644)
}
