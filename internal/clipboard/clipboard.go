package clipboard

// Clipboard interface defines methods for clipboard operations
type Clipboard interface {
	ReadImage() ([]byte, string, error)
}
