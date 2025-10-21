//go:build windows
// +build windows

package screenshot

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

// CaptureScreen captures the entire primary screen and saves it to a file
func CaptureScreen(outputPath string) error {
	// Use default path if none specified
	if outputPath == "" {
		timestamp := time.Now().Format("20060102_150405")
		outputPath = GetScreenshotPath(fmt.Sprintf("screenshot_%s.png", timestamp))
	}
	// Get screen dimensions
	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	// Get device contexts
	hdcScreen := win.GetDC(0)
	defer win.ReleaseDC(0, hdcScreen)

	hdcMem := win.CreateCompatibleDC(hdcScreen)
	defer win.DeleteDC(hdcMem)

	// Create bitmap
	bitmapInfo := win.BITMAPINFO{
		BmiHeader: win.BITMAPINFOHEADER{
			BiSize:        uint32(unsafe.Sizeof(win.BITMAPINFOHEADER{})),
			BiWidth:       int32(screenWidth),
			BiHeight:      -int32(screenHeight), // Negative for top-down DIB
			BiPlanes:      1,
			BiBitCount:    32,
			BiCompression: win.BI_RGB,
		},
	}
	var ppvBits unsafe.Pointer
	hbmp := win.CreateDIBSection(
		hdcMem,
		&bitmapInfo.BmiHeader,
		win.DIB_RGB_COLORS,
		&ppvBits,
		0,
		0,
	)
	defer win.DeleteObject(win.HGDIOBJ(hbmp))

	// Select bitmap into memory DC
	oldBmp := win.SelectObject(hdcMem, win.HGDIOBJ(hbmp))
	defer win.SelectObject(hdcMem, oldBmp)

	// Copy screen content
	win.BitBlt(
		hdcMem,
		0, 0,
		int32(screenWidth), int32(screenHeight),
		hdcScreen,
		0, 0,
		win.SRCCOPY,
	)

	// Prepare bitmap data for PNG encoding
	bmpHeader := struct {
		Signature [2]byte
		FileSize  uint32
		Reserved  uint32
		DataOff   uint32
	}{
		[2]byte{'B', 'M'},
		0,
		0,
		54,
	}

	infoHeaderSize := uint32(40)
	fileHeaderSize := uint32(14)
	imageSize := uint32(screenWidth * screenHeight * 4) // 4 bytes per pixel (BGRA)
	totalSize := fileHeaderSize + infoHeaderSize + imageSize

	bmpHeader.FileSize = totalSize
	bmpHeader.DataOff = fileHeaderSize + infoHeaderSize

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, bmpHeader)
	binary.Write(buf, binary.LittleEndian, bitmapInfo.BmiHeader)
	binary.Write(buf, binary.LittleEndian, (*[1 << 30]byte)(ppvBits)[:imageSize])

	// Convert to image and save as PNG
	img, err := bmpToImage(buf.Bytes(), screenWidth, screenHeight)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// GetScreenDimensions returns the width and height of the primary screen
func GetScreenDimensions() (int, int) {
	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	return screenWidth, screenHeight
}

// Helper function to convert BMP bytes to image.RGBA
func bmpToImage(bmpData []byte, width, height int) (*image.RGBA, error) {
	// BMP data starts at offset 54 (headers)
	const headerSize = 54
	if len(bmpData) < headerSize {
		return nil, errors.New("invalid BMP data")
	}

	// Create RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// BMP is stored bottom-to-top, so flip vertically
	bytesPerPixel := 4
	rowSize := width * bytesPerPixel
	for y := 0; y < height; y++ {
		srcY := height - 1 - y
		srcStart := headerSize + srcY*rowSize
		dstStart := y * img.Stride
		copy(
			img.Pix[dstStart:dstStart+rowSize],
			bmpData[srcStart:srcStart+rowSize],
		)
	}
	return img, nil
}
