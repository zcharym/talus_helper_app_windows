package services

import (
	"context"

	"talus_helper_windows/internal/config"
)

// ConfigService handles configuration-related operations
type ConfigService struct {
	ctx    context.Context
	config *config.Config
}

// NewConfigService creates a new ConfigService
func NewConfigService(ctx context.Context, cfg *config.Config) *ConfigService {
	return &ConfigService{
		ctx:    ctx,
		config: cfg,
	}
}

// GetConfig returns the current configuration
func (s *ConfigService) GetConfig() (config.Config, error) {
	if s.config == nil {
		return config.GetDefault(), nil
	}
	return *s.config, nil
}

// SaveConfig saves the configuration
func (s *ConfigService) SaveConfig(cfg config.Config) error {
	if err := config.Save(cfg); err != nil {
		return err
	}
	// Update the in-memory config
	s.config = &cfg
	return nil
}
