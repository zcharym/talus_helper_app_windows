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
	return s.storage.LoadTodos()
}

// AddTodo adds a new todo
func (s *TodoService) AddTodo(text string) (models.Todo, error) {
	todos, err := s.storage.LoadTodos()
	if err != nil {
		return models.Todo{}, err
	}

	newTodo := models.Todo{
		ID:        uuid.New().String(),
		Text:      text,
		Completed: false,
		CreatedAt: time.Now(),
	}

	todos = append(todos, newTodo)

	if err := s.storage.SaveTodos(todos); err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

// UpdateTodo updates an existing todo
func (s *TodoService) UpdateTodo(id, text string, completed bool) (models.Todo, error) {
	todos, err := s.storage.LoadTodos()
	if err != nil {
		return models.Todo{}, err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Text = text
			todos[i].Completed = completed
			if err := s.storage.SaveTodos(todos); err != nil {
				return models.Todo{}, err
			}
			return todos[i], nil
		}
	}

	return models.Todo{}, fmt.Errorf("todo with id %s not found", id)
}

// DeleteTodo deletes a todo
func (s *TodoService) DeleteTodo(id string) error {
	todos, err := s.storage.LoadTodos()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return s.storage.SaveTodos(todos)
		}
	}

	return fmt.Errorf("todo with id %s not found", id)
}
