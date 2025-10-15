package clipboard

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"golang.design/x/clipboard"
)

// WindowsClipboard implements Clipboard interface for Windows
type WindowsClipboard struct{}

// NewWindowsClipboard creates a new Windows clipboard instance
func NewWindowsClipboard() *WindowsClipboard {
	return &WindowsClipboard{}
}

// ReadImage reads image data from the Windows clipboard
func (w *WindowsClipboard) ReadImage() ([]byte, string, error) {
	// Read clipboard content
	clipboardData := clipboard.Read(clipboard.FmtImage)
	if len(clipboardData) == 0 {
		return nil, "", fmt.Errorf("no image found in clipboard")
	}

	// Detect image format
	format := detectFormat(clipboardData)
	if format == "" {
		return nil, "", fmt.Errorf("unsupported image format")
	}

	return clipboardData, format, nil
}

// detectFormat detects the image format from the data
func detectFormat(data []byte) string {
	if len(data) < 8 {
		return ""
	}

	// Check for PNG signature
	if len(data) >= 8 && data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "png"
	}

	// Check for JPEG signature
	if len(data) >= 2 && data[0] == 0xFF && data[1] == 0xD8 {
		return "jpeg"
	}

	// Check for BMP signature
	if len(data) >= 2 && data[0] == 0x42 && data[1] == 0x4D {
		return "bmp"
	}

	// Try to decode as image to validate
	_, format, err := image.DecodeConfig(io.Reader(bytes.NewReader(data)))
	if err != nil {
		return ""
	}

	return format
}
