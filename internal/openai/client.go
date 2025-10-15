package openai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the OpenAI API client
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewClient creates a new OpenAI client
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ExtractTextFromImage extracts text from an image using Vision API
func (c *Client) ExtractTextFromImage(imageData []byte, imageFormat string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("API key is required")
	}

	// Encode image to base64
	base64Image := encodeToBase64(imageData)

	// Build the request
	request := buildVisionRequest(base64Image, imageFormat)

	// Marshal request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.BaseURL+"/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		var errorResp ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResp); err != nil {
			return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(responseBody))
		}
		return "", fmt.Errorf("API error: %s", errorResp.Error.Message)
	}

	// Parse response
	var visionResp VisionResponse
	if err := json.Unmarshal(responseBody, &visionResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Extract text from response
	if len(visionResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	// Get the text content from the first choice
	choice := visionResp.Choices[0]
	if textContent, ok := choice.Message.Content.(string); ok {
		return textContent, nil
	}

	return "", fmt.Errorf("unexpected response format")
}

// encodeToBase64 encodes data to base64 string
func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// buildVisionRequest builds the Vision API request
func buildVisionRequest(base64Image, format string) *VisionRequest {
	// Determine MIME type based on format
	mimeType := "image/png"
	switch format {
	case "jpeg", "jpg":
		mimeType = "image/jpeg"
	case "bmp":
		mimeType = "image/bmp"
	}

	// Create the data URL
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image)

	return &VisionRequest{
		Model: "moonshot-v1-8k-vision-preview",
		Messages: []Message{
			{
				Role: "system",
				Content: TextContent{
					Type: "text",
					Text: "Extract all text from images accurately. Return only the text content, no explanations or additional formatting.",
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ImageContent{
						Type: "image_url",
						ImageURL: ImageURL{
							URL: dataURL,
						},
					},
					TextContent{
						Type: "text",
						Text: "Extract all text from this image. Return only the text content.",
					},
				},
			},
		},
		Temperature: 0.3,
	}
}
