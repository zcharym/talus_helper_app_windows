package encode

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

// Protocol constants
const (
	// Header size in bytes (metadata)
	HeaderSize = 32
	// Channels per pixel (RGB)
	ChannelsPerPixel = 3
	// Bits per channel (8 bits = 1 byte)
	BitsPerChannel = 8
	// Bytes per pixel
	BytesPerPixel = ChannelsPerPixel
	// Magic number for protocol identification
	MagicNumber uint32 = 0x54414C55 // "TALU" (changed from TIMG to match project)
	// Version
	Version uint16 = 1
)

// Header structure for the encoded data
type Header struct {
	Magic          uint32   // Magic number for identification
	Version        uint16   // Protocol version
	Reserved       uint16   // Reserved for future use
	OriginalSize   uint32   // Size of original uncompressed data
	CompressedSize uint32   // Size of compressed data
	Checksum       [16]byte // SHA256 checksum (first 16 bytes)
}

// Encoder handles text-to-image conversion
type Encoder struct {
	compression bool
}

// Decoder handles image-to-text conversion
type Decoder struct{}

// NewEncoder creates a new encoder instance
func NewEncoder(useCompression bool) *Encoder {
	return &Encoder{compression: useCompression}
}

// NewDecoder creates a new decoder instance
func NewDecoder() *Decoder {
	return &Decoder{}
}

// EncodeTextToImage converts text data to PNG image
func (e *Encoder) EncodeTextToImage(data []byte, outputPath string) error {
	// Validate input
	if len(data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}

	// Compress data if enabled
	var processedData []byte
	var err error

	if e.compression {
		processedData, err = e.compressData(data)
		if err != nil {
			return fmt.Errorf("compression failed: %w", err)
		}
	} else {
		processedData = data
	}

	// Create header
	header := e.createHeader(data, processedData)

	// Serialize header
	headerBytes, err := e.serializeHeader(header)
	if err != nil {
		return fmt.Errorf("header serialization failed: %w", err)
	}

	// Combine header and data
	fullData := append(headerBytes, processedData...)

	// Calculate optimal image dimensions
	width, height := e.calculateDimensions(len(fullData))

	// Create image
	img := e.createImage(fullData, width, height)

	// Save image
	return e.saveImage(img, outputPath)
}

// DecodeTextFromImage converts PNG image back to text data
func (d *Decoder) DecodeTextFromImage(imagePath string) ([]byte, error) {
	// Load image
	img, err := d.loadImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %w", err)
	}

	// Extract data from image
	data := d.extractData(img)

	// Parse header
	if len(data) < HeaderSize {
		return nil, fmt.Errorf("invalid image: too small to contain header")
	}

	header, err := d.parseHeader(data[:HeaderSize])
	if err != nil {
		return nil, fmt.Errorf("failed to parse header: %w", err)
	}

	// Validate header
	if err := d.validateHeader(header); err != nil {
		return nil, fmt.Errorf("invalid header: %w", err)
	}

	// Extract payload
	payloadSize := int(header.CompressedSize)
	if len(data) < HeaderSize+payloadSize {
		return nil, fmt.Errorf("invalid image: insufficient data")
	}

	payload := data[HeaderSize : HeaderSize+payloadSize]

	// Decompress if needed
	var originalData []byte
	if header.CompressedSize != header.OriginalSize {
		originalData, err = d.decompressData(payload)
		if err != nil {
			return nil, fmt.Errorf("decompression failed: %w", err)
		}
	} else {
		originalData = payload
	}

	// Verify checksum
	if !d.verifyChecksum(originalData, header.Checksum) {
		return nil, fmt.Errorf("checksum verification failed")
	}

	return originalData, nil
}

// compressData compresses the input data using gzip
func (e *Encoder) compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	if _, err := writer.Write(data); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// decompressData decompresses gzip data
func (d *Decoder) decompressData(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

// createHeader creates a header for the encoded data
func (e *Encoder) createHeader(original, compressed []byte) Header {
	hash := sha256.Sum256(original)
	var checksum [16]byte
	copy(checksum[:], hash[:16])

	return Header{
		Magic:          MagicNumber,
		Version:        Version,
		Reserved:       0,
		OriginalSize:   uint32(len(original)),
		CompressedSize: uint32(len(compressed)),
		Checksum:       checksum,
	}
}

// serializeHeader converts header to bytes
func (e *Encoder) serializeHeader(header Header) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, header); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// parseHeader parses header from bytes
func (d *Decoder) parseHeader(data []byte) (Header, error) {
	var header Header
	reader := bytes.NewReader(data)

	if err := binary.Read(reader, binary.BigEndian, &header); err != nil {
		return Header{}, err
	}

	return header, nil
}

// validateHeader validates header fields
func (d *Decoder) validateHeader(header Header) error {
	if header.Magic != MagicNumber {
		return fmt.Errorf("invalid magic number: %x", header.Magic)
	}

	if header.Version != Version {
		return fmt.Errorf("unsupported version: %d", header.Version)
	}

	return nil
}

// verifyChecksum verifies data integrity
func (d *Decoder) verifyChecksum(data []byte, expectedChecksum [16]byte) bool {
	hash := sha256.Sum256(data)
	return bytes.Equal(hash[:16], expectedChecksum[:])
}

// calculateDimensions calculates optimal image dimensions
func (e *Encoder) calculateDimensions(dataSize int) (int, int) {
	// Calculate minimum pixels needed
	pixelsNeeded := int(math.Ceil(float64(dataSize) / float64(BytesPerPixel)))

	// Calculate square-ish dimensions
	width := int(math.Ceil(math.Sqrt(float64(pixelsNeeded))))
	height := int(math.Ceil(float64(pixelsNeeded) / float64(width)))

	return width, height
}

// createImage creates an image from data
func (e *Encoder) createImage(data []byte, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	dataIndex := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b uint8 = 0, 0, 0

			// Pack 3 bytes into RGB channels
			if dataIndex < len(data) {
				r = data[dataIndex]
				dataIndex++
			}
			if dataIndex < len(data) {
				g = data[dataIndex]
				dataIndex++
			}
			if dataIndex < len(data) {
				b = data[dataIndex]
				dataIndex++
			}

			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	return img
}

// saveImage saves image to file
func (e *Encoder) saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// loadImage loads image from file
func (d *Decoder) loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

// extractData extracts data from image
func (d *Decoder) extractData(img image.Image) []byte {
	bounds := img.Bounds()
	var data []byte

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Extract bytes from RGB channels
			data = append(data, uint8(r>>8))
			data = append(data, uint8(g>>8))
			data = append(data, uint8(b>>8))
		}
	}

	return data
}
