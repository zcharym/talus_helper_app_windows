package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents application configuration
type Config struct {
	Theme               string `json:"theme"`
	AutoSave            bool   `json:"autoSave"`
	Notifications       bool   `json:"notifications"`
	OpenAIAPIKey        string `json:"openAIAPIKey"`
	OpenAIBaseURL       string `json:"openAIBaseURL"`
	DefaultTodoCategory string `json:"defaultTodoCategory"`
	MaxTodos            int    `json:"maxTodos"`
	Language            string `json:"language"`
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
	}
}

// Load reads the configuration from file
func Load() (*Config, error) {
	dataDir, err := GetDataDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(dataDir, "config.json")
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
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Save writes the configuration to file
func Save(config Config) error {
	dataDir, err := GetDataDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(dataDir, "config.json")
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)
}
