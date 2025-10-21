package main

import (
	"context"
	"fmt"

	"talus_helper_windows/internal/clipboard"
	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/models"
	"talus_helper_windows/internal/services"
	"talus_helper_windows/internal/storage"
)

// App struct - thin orchestration layer
type App struct {
	ctx              context.Context
	config           *config.Config
	storage          storage.Storage
	clipboard        clipboard.Clipboard
	todoService      *services.TodoService
	configService    *services.ConfigService
	clipboardService *services.ClipboardService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load environment variables for debug mode
	config.LoadEnvForDebug()

	// Initialize dependencies
	var err error
	a.config, err = config.Load()
	if err != nil {
		// Use default config if loading fails
		defaultConfig := config.GetDefault()
		a.config = &defaultConfig
	}

	// Initialize SQLite storage
	a.storage = storage.NewSQLiteStorage()
	if err := a.storage.Connect(ctx); err != nil {
		// Log error but continue with default config
		// In production, you might want to handle this more gracefully
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	// Run database migrations
	if err := a.storage.Migrate(ctx); err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
	}

	a.clipboard = clipboard.NewWindowsClipboard()

	// Initialize services
	a.todoService = services.NewTodoService(ctx, a.storage)
	a.configService = services.NewConfigService(ctx, a.config)
	a.clipboardService = services.NewClipboardService(ctx, a.config, a.clipboard)
}

// Todo methods - delegated to TodoService

// GetTodos returns all todos
func (a *App) GetTodos() ([]models.Todo, error) {
	return a.todoService.GetTodos()
}

// AddTodo adds a new todo
func (a *App) AddTodo(text string) (models.Todo, error) {
	return a.todoService.AddTodo(text)
}

// UpdateTodo updates an existing todo
func (a *App) UpdateTodo(id, text string, completed bool) (models.Todo, error) {
	return a.todoService.UpdateTodo(id, text, completed)
}

// DeleteTodo deletes a todo
func (a *App) DeleteTodo(id string) error {
	return a.todoService.DeleteTodo(id)
}

// Config methods - delegated to ConfigService

// GetConfig returns the current configuration
func (a *App) GetConfig() (config.Config, error) {
	return a.configService.GetConfig()
}

// SaveConfig saves the configuration
func (a *App) SaveConfig(cfg config.Config) error {
	return a.configService.SaveConfig(cfg)
}

// Clipboard methods - delegated to ClipboardService

// OCRFromClipboard extracts text from clipboard image using OpenAI Vision API
func (a *App) OCRFromClipboard() (string, error) {
	return a.clipboardService.OCRFromClipboard()
}

// shutdown is called when the app shuts down
func (a *App) shutdown(ctx context.Context) {
	if a.storage != nil {
		if err := a.storage.Close(); err != nil {
			fmt.Printf("Failed to close database connection: %v\n", err)
		}
	}
}
