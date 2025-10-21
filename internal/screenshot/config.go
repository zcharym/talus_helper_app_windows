package screenshot

import (
	"os"
	"path/filepath"
)

// DefaultOutputPath is the default directory for saving screenshots
const DefaultOutputPath = "data/screenshots"

// GetOutputPath returns the output path, creating the directory if it doesn't exist
func GetOutputPath() string {
	// Ensure the directory exists
	if err := os.MkdirAll(DefaultOutputPath, 0755); err != nil {
		// If we can't create the directory, fall back to current directory
		return "screenshots"
	}
	return DefaultOutputPath
}

// GetScreenshotPath returns a full path for a screenshot file
func GetScreenshotPath(filename string) string {
	outputDir := GetOutputPath()
	return filepath.Join(outputDir, filename)
}
