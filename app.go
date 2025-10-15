package main

import (
	"context"
	"fmt"
	"time"

	"talus_helper_windows/internal/clipboard"
	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/models"
	"talus_helper_windows/internal/openai"
	"talus_helper_windows/internal/storage"

	"github.com/google/uuid"
)

// App struct - thin orchestration layer
type App struct {
	ctx          context.Context
	config       *config.Config
	storage      storage.Storage
	clipboard    clipboard.Clipboard
	openaiClient *openai.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Initialize dependencies
	var err error
	a.config, err = config.Load()
	if err != nil {
		// Use default config if loading fails
		defaultConfig := config.GetDefault()
		a.config = &defaultConfig
	}

	a.storage = storage.NewJSONStorage()
	a.clipboard = clipboard.NewWindowsClipboard()
	// openaiClient will be initialized on-demand in OCRFromClipboard
}

// Todo methods

// GetTodos returns all todos
func (a *App) GetTodos() ([]models.Todo, error) {
	return a.storage.LoadTodos()
}

// AddTodo adds a new todo
func (a *App) AddTodo(text string) (models.Todo, error) {
	todos, err := a.storage.LoadTodos()
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

	if err := a.storage.SaveTodos(todos); err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

// UpdateTodo updates an existing todo
func (a *App) UpdateTodo(id, text string, completed bool) (models.Todo, error) {
	todos, err := a.storage.LoadTodos()
	if err != nil {
		return models.Todo{}, err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Text = text
			todos[i].Completed = completed
			if err := a.storage.SaveTodos(todos); err != nil {
				return models.Todo{}, err
			}
			return todos[i], nil
		}
	}

	return models.Todo{}, fmt.Errorf("todo with id %s not found", id)
}

// DeleteTodo deletes a todo
func (a *App) DeleteTodo(id string) error {
	todos, err := a.storage.LoadTodos()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return a.storage.SaveTodos(todos)
		}
	}

	return fmt.Errorf("todo with id %s not found", id)
}

// Config methods

// GetConfig returns the current configuration
func (a *App) GetConfig() (config.Config, error) {
	if a.config == nil {
		return config.GetDefault(), nil
	}
	return *a.config, nil
}

// SaveConfig saves the configuration
func (a *App) SaveConfig(cfg config.Config) error {
	if err := config.Save(cfg); err != nil {
		return err
	}
	// Update the in-memory config
	a.config = &cfg
	return nil
}

// OCRFromClipboard extracts text from clipboard image using OpenAI Vision API
func (a *App) OCRFromClipboard() (string, error) {
	// Validate API key and base URL
	if a.config.OpenAIAPIKey == "" {
		return "", fmt.Errorf("OpenAI API key is not configured. Please set it in Settings")
	}
	if a.config.OpenAIBaseURL == "" {
		return "", fmt.Errorf("OpenAI Base URL is not configured. Please set it in Settings")
	}

	// Read image from clipboard
	imageData, format, err := a.clipboard.ReadImage()
	if err != nil {
		return "", fmt.Errorf("failed to read image from clipboard: %w", err)
	}

	// Initialize OpenAI client if not already done
	if a.openaiClient == nil {
		a.openaiClient = openai.NewClient(a.config.OpenAIBaseURL, a.config.OpenAIAPIKey)
	}

	// Extract text from image
	text, err := a.openaiClient.ExtractTextFromImage(imageData, format)
	if err != nil {
		return "", fmt.Errorf("failed to extract text from image: %w", err)
	}

	return text, nil
}
