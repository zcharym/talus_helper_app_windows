package services

import (
	"context"
	"fmt"
	"time"

	"talus_helper_windows/internal/models"
	"talus_helper_windows/internal/storage"

	"github.com/google/uuid"
)

// TodoService handles todo-related operations
type TodoService struct {
	ctx     context.Context
	storage storage.Storage
}

// NewTodoService creates a new TodoService
func NewTodoService(ctx context.Context, storage storage.Storage) *TodoService {
	return &TodoService{
		ctx:     ctx,
		storage: storage,
	}
}

// GetTodos returns all todos
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.storage.GetTodos(s.ctx)
}

// AddTodo adds a new todo
func (s *TodoService) AddTodo(text string) (models.Todo, error) {
	newTodo := models.Todo{
		ID:        uuid.New().String(),
		Text:      text,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := s.storage.CreateTodo(s.ctx, &newTodo); err != nil {
		return models.Todo{}, fmt.Errorf("failed to create todo: %w", err)
	}

	return newTodo, nil
}

// UpdateTodo updates an existing todo
func (s *TodoService) UpdateTodo(id, text string, completed bool) (models.Todo, error) {
	// First, get the existing todo to preserve the created_at timestamp
	existingTodo, err := s.storage.GetTodoByID(s.ctx, id)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to get todo: %w", err)
	}

	// Update the todo with new values
	existingTodo.Text = text
	existingTodo.Completed = completed

	if err := s.storage.UpdateTodo(s.ctx, existingTodo); err != nil {
		return models.Todo{}, fmt.Errorf("failed to update todo: %w", err)
	}

	return *existingTodo, nil
}

// DeleteTodo deletes a todo
func (s *TodoService) DeleteTodo(id string) error {
	if err := s.storage.DeleteTodo(s.ctx, id); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}
