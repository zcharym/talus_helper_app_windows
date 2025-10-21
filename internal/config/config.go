package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

// Config represents application configuration
type Config struct {
	Theme               string `toml:"theme"`
	AutoSave            bool   `toml:"autoSave"`
	Notifications       bool   `toml:"notifications"`
	OpenAIBaseURL       string `toml:"openAIBaseURL"`
	DefaultTodoCategory string `toml:"defaultTodoCategory"`
	MaxTodos            int    `toml:"maxTodos"`
	Language            string `toml:"language"`
	OpenAIAPIKey        string `toml:"openAIAPIKey"`
	WorkflowyAPIKey     string `toml:"workflowyAPIKey"`
}

// LoadEnvForDebug loads .env file if it exists (for debug mode)
func LoadEnvForDebug() {
	// Try to load .env file from current directory
	if err := godotenv.Load(); err != nil {
		// .env file doesn't exist or can't be read, that's okay
		fmt.Printf("Debug: No .env file found or error loading: %v\n", err)
	} else {
		fmt.Println("Debug: Loaded environment variables from .env file")
	}
}

// GetDataDir returns the application data directory
func GetDataDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dataDir := filepath.Join(homeDir, ".talus-helper")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return "", err
	}
	return dataDir, nil
}

// GetDefault returns the default configuration
func GetDefault() Config {
	return Config{
		Theme:               "light",
		AutoSave:            true,
		Notifications:       true,
		OpenAIAPIKey:        "",
		OpenAIBaseURL:       "https://api.moonshot.cn/v1",
		DefaultTodoCategory: "General",
		MaxTodos:            100,
		Language:            "en",
		WorkflowyAPIKey:     "",
	}
}

// Load reads the configuration from file and environment variables
func Load() (*Config, error) {
	dataDir, err := GetDataDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(dataDir, "config.toml")
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default config if file doesn't exist
			defaultConfig := GetDefault()
			return &defaultConfig, nil
		}
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// Override with environment variables if they exist
	if envAPIKey := os.Getenv("OPENAI_API_KEY"); envAPIKey != "" {
		config.OpenAIAPIKey = envAPIKey
	}
	if envBaseURL := os.Getenv("OPENAI_BASE_URL"); envBaseURL != "" {
		config.OpenAIBaseURL = envBaseURL
	}
	if envTheme := os.Getenv("THEME"); envTheme != "" {
		config.Theme = envTheme
	}
	if envLanguage := os.Getenv("LANGUAGE"); envLanguage != "" {
		config.Language = envLanguage
	}
	if envWorkflowyAPIKey := os.Getenv("WORKFLOWY_API_KEY"); envWorkflowyAPIKey != "" {
		config.WorkflowyAPIKey = envWorkflowyAPIKey
	}

	return &config, nil
}

// Save writes the configuration to file
func Save(config Config) error {
	dataDir, err := GetDataDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(dataDir, "config.toml")
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	return encoder.Encode(config)
}
