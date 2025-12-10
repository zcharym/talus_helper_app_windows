package services

import (
	"context"
	"fmt"

	"talus_helper_windows/internal/config"
	"talus_helper_windows/internal/openai"
)

// PromptOptimizerService handles prompt optimization operations
type PromptOptimizerService struct {
	ctx          context.Context
	config       *config.Config
	openaiClient *openai.Client
}

// NewPromptOptimizerService creates a new PromptOptimizerService
func NewPromptOptimizerService(ctx context.Context, cfg *config.Config) *PromptOptimizerService {
	return &PromptOptimizerService{
		ctx:    ctx,
		config: cfg,
	}
}

// getOptimizationPrompt returns the system prompt based on the optimization style
func getOptimizationPrompt(style string) string {
	switch style {
	case "clarity":
		return "You are a prompt optimization expert. Improve the clarity and readability of the given prompt. Make it easier to understand while preserving its original intent and meaning. Return only the optimized prompt, no explanations."
	case "specificity":
		return "You are a prompt optimization expert. Enhance the specificity of the given prompt by adding more concrete details, examples, and precise requirements. Make it more actionable while preserving its original intent. Return only the optimized prompt, no explanations."
	case "ai-focused":
		return "You are a prompt optimization expert specializing in AI/LLM interactions. Optimize the given prompt to be more effective for AI language models. Use clear structure, explicit instructions, and format requirements. Return only the optimized prompt, no explanations."
	case "concise":
		return "You are a prompt optimization expert. Make the given prompt more concise and to-the-point while preserving all essential information and intent. Remove redundancy and unnecessary words. Return only the optimized prompt, no explanations."
	case "detailed":
		return "You are a prompt optimization expert. Expand the given prompt with more detail, context, and comprehensive information while maintaining clarity. Add relevant examples and specifications. Return only the optimized prompt, no explanations."
	default:
		return "You are a prompt optimization expert. Improve the given prompt to make it clearer, more specific, and more effective. Preserve its original intent and meaning. Return only the optimized prompt, no explanations."
	}
}

// OptimizePrompt optimizes a user prompt using OpenAI API
func (s *PromptOptimizerService) OptimizePrompt(prompt string, style string) (string, error) {
	// Validate API key and base URL
	if s.config.OpenAIAPIKey == "" {
		return "", fmt.Errorf("OpenAI API key is not configured. Please set it in Settings")
	}
	if s.config.OpenAIBaseURL == "" {
		return "", fmt.Errorf("OpenAI Base URL is not configured. Please set it in Settings")
	}

	// Validate input
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	// Initialize OpenAI client if not already done
	if s.openaiClient == nil {
		s.openaiClient = openai.NewClient(s.config.OpenAIBaseURL, s.config.OpenAIAPIKey)
	}

	// Get optimization prompt based on style
	systemPrompt := getOptimizationPrompt(style)

	// Build messages for chat completion
	// Use strings directly for text content (OpenAI standard format)
	messages := []openai.Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// Call chat completion API
	optimizedPrompt, err := s.openaiClient.ChatCompletion(messages, 0.7)
	if err != nil {
		return "", fmt.Errorf("failed to optimize prompt: %w", err)
	}

	return optimizedPrompt, nil
}

