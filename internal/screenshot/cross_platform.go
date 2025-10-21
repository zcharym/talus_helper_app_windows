package screenshot

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

// DisplayInfo represents information about a display
type DisplayInfo struct {
	Index  int
	Bounds image.Rectangle
}

// CaptureScreen captures the entire primary screen and saves it to a file
func CaptureScreen(outputPath string) error {
	// Use default path if none specified
	if outputPath == "" {
		timestamp := time.Now().Format("20060102_150405")
		outputPath = GetScreenshotPath(fmt.Sprintf("screenshot_%s.png", timestamp))
	}

	// Capture primary display (display 0)
	return CaptureDisplay(0, outputPath)
}

// GetScreenDimensions returns the width and height of the primary screen
func GetScreenDimensions() (int, int) {
	bounds := screenshot.GetDisplayBounds(0)
	return bounds.Dx(), bounds.Dy()
}

// ListDisplays returns information about all active displays
func ListDisplays() ([]DisplayInfo, error) {
	n := screenshot.NumActiveDisplays()
	displays := make([]DisplayInfo, n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		displays[i] = DisplayInfo{
			Index:  i,
			Bounds: bounds,
		}
	}

	return displays, nil
}

// CaptureDisplay captures a specific display and saves it to a file
func CaptureDisplay(displayIndex int, outputPath string) error {
	n := screenshot.NumActiveDisplays()
	if displayIndex < 0 || displayIndex >= n {
		return fmt.Errorf("invalid display index %d, available displays: 0-%d", displayIndex, n-1)
	}

	bounds := screenshot.GetDisplayBounds(displayIndex)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return fmt.Errorf("failed to capture display %d: %w", displayIndex, err)
	}

	// Use default path if none specified
	if outputPath == "" {
		timestamp := time.Now().Format("20060102_150405")
		outputPath = GetScreenshotPath(fmt.Sprintf("display_%d_%s.png", displayIndex, timestamp))
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	return png.Encode(file, img)
}

// CaptureAllDisplays captures all displays and saves them to files
func CaptureAllDisplays(outputDir string) error {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return fmt.Errorf("failed to capture display %d: %w", i, err)
		}

		fileName := fmt.Sprintf("%s/display_%d_%dx%d.png", outputDir, i, bounds.Dx(), bounds.Dy())
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", fileName, err)
		}

		err = png.Encode(file, img)
		file.Close()
		if err != nil {
			return fmt.Errorf("failed to encode image for display %d: %w", i, err)
		}
	}

	return nil
}
