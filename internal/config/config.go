package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config represents application configuration
type Config struct {
	Theme               string `toml:"theme"`
	AutoSave            bool   `toml:"autoSave"`
	Notifications       bool   `toml:"notifications"`
	OpenAIAPIKey        string `toml:"openAIAPIKey"`
	OpenAIBaseURL       string `toml:"openAIBaseURL"`
	DefaultTodoCategory string `toml:"defaultTodoCategory"`
	MaxTodos            int    `toml:"maxTodos"`
	Language            string `toml:"language"`
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
