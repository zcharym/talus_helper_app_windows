package services

import (
	"context"
	"fmt"

	"talus_helper_windows/internal/clipboard"
	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/openai"
)

// ClipboardService handles clipboard and OCR operations
type ClipboardService struct {
	ctx           context.Context
	config        *config.Config
	clipboard     clipboard.Clipboard
	openaiClient  *openai.Client
}

// NewClipboardService creates a new ClipboardService
func NewClipboardService(ctx context.Context, cfg *config.Config, clipboard clipboard.Clipboard) *ClipboardService {
	return &ClipboardService{
		ctx:       ctx,
		config:    cfg,
		clipboard: clipboard,
	}
}

// OCRFromClipboard extracts text from clipboard image using OpenAI Vision API
func (s *ClipboardService) OCRFromClipboard() (string, error) {
	// Validate API key and base URL
	if s.config.OpenAIAPIKey == "" {
		return "", fmt.Errorf("OpenAI API key is not configured. Please set it in Settings")
	}
	if s.config.OpenAIBaseURL == "" {
		return "", fmt.Errorf("OpenAI Base URL is not configured. Please set it in Settings")
	}

	// Read image from clipboard
	imageData, format, err := s.clipboard.ReadImage()
	if err != nil {
		return "", fmt.Errorf("failed to read image from clipboard: %w", err)
	}

	// Initialize OpenAI client if not already done
	if s.openaiClient == nil {
		s.openaiClient = openai.NewClient(s.config.OpenAIBaseURL, s.config.OpenAIAPIKey)
	}

	// Extract text from image
	text, err := s.openaiClient.ExtractTextFromImage(imageData, format)
	if err != nil {
		return "", fmt.Errorf("failed to extract text from image: %w", err)
	}

	return text, nil
}
