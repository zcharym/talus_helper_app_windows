package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/models"

	_ "modernc.org/sqlite"
)

// Storage interface defines methods for data persistence
type Storage interface {
	// Connection management
	Connect(ctx context.Context) error
	Close() error

	// Todo operations
	GetTodos(ctx context.Context) ([]models.Todo, error)
	GetTodoByID(ctx context.Context, id string) (*models.Todo, error)
	CreateTodo(ctx context.Context, todo *models.Todo) error
	UpdateTodo(ctx context.Context, todo *models.Todo) error
	DeleteTodo(ctx context.Context, id string) error

	// Database management
	Migrate(ctx context.Context) error
}

// SQLiteStorage implements Storage interface using SQLite database
type SQLiteStorage struct {
	db      *sql.DB
	dataDir string
}

// NewSQLiteStorage creates a new SQLite storage instance
func NewSQLiteStorage() *SQLiteStorage {
	return &SQLiteStorage{}
}

// Connect establishes connection to the SQLite database
func (s *SQLiteStorage) Connect(ctx context.Context) error {
	var err error
	s.dataDir, err = config.GetDataDir()
	if err != nil {
		return fmt.Errorf("failed to get data directory: %w", err)
	}

	// Ensure data directory exists
	if err := os.MkdirAll(s.dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	dbPath := filepath.Join(s.dataDir, "todos.db")
	s.db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// Close closes the database connection
func (s *SQLiteStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// Migrate creates the necessary database tables
func (s *SQLiteStorage) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id TEXT PRIMARY KEY,
		text TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_todos_created_at ON todos(created_at);
	CREATE INDEX IF NOT EXISTS idx_todos_completed ON todos(completed);
	`

	_, err := s.db.ExecContext(ctx, query)
	return err
}

// GetTodos retrieves all todos from the database
func (s *SQLiteStorage) GetTodos(ctx context.Context) ([]models.Todo, error) {
	query := `SELECT id, text, completed, created_at FROM todos ORDER BY created_at DESC`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Text, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return todos, nil
}

// GetTodoByID retrieves a specific todo by ID
func (s *SQLiteStorage) GetTodoByID(ctx context.Context, id string) (*models.Todo, error) {
	query := `SELECT id, text, completed, created_at FROM todos WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, id)

	var todo models.Todo
	err := row.Scan(&todo.ID, &todo.Text, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("todo with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to scan todo: %w", err)
	}

	return &todo, nil
}

// CreateTodo creates a new todo in the database
func (s *SQLiteStorage) CreateTodo(ctx context.Context, todo *models.Todo) error {
	query := `INSERT INTO todos (id, text, completed, created_at) VALUES (?, ?, ?, ?)`
	_, err := s.db.ExecContext(ctx, query, todo.ID, todo.Text, todo.Completed, todo.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}
	return nil
}

// UpdateTodo updates an existing todo in the database
func (s *SQLiteStorage) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	query := `UPDATE todos SET text = ?, completed = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, todo.Text, todo.Completed, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %s not found", todo.ID)
	}

	return nil
}

// DeleteTodo deletes a todo from the database
func (s *SQLiteStorage) DeleteTodo(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %s not found", id)
	}

	return nil
}
