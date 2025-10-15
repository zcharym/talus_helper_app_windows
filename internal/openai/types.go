package openai

// VisionRequest represents the request structure for Moonshot Vision API
type VisionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// Message represents a message in the conversation
type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

// TextContent represents text content in a message
type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ImageContent represents image content in a message
type ImageContent struct {
	Type     string    `json:"type"`
	ImageURL ImageURL  `json:"image_url"`
}

// ImageURL represents the image URL structure
type ImageURL struct {
	URL string `json:"url"`
}

// VisionResponse represents the response from Moonshot Vision API
type VisionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice represents a choice in the response
type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
	FinishReason string `json:"finish_reason"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error APIError `json:"error"`
}

// APIError represents API error details
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}
