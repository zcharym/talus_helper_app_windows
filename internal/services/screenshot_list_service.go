package services

import (
    "errors"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "sort"
    "strings"
    "time"
)

// ScreenshotFile represents basic metadata about a screenshot file
type ScreenshotFile struct {
    Name    string    `json:"name"`
    Path    string    `json:"path"`
    Size    int64     `json:"size"`
    ModTime time.Time `json:"modTime"`
}

// ScreenshotListService lists and opens screenshots from a base directory
type ScreenshotListService struct {
    BaseDir string
}

// NewScreenshotListService constructs a new service using the default
// screenshots directory at ~/.talus-helper/screenshots
func NewScreenshotListService() *ScreenshotListService {
    home, _ := os.UserHomeDir()
    base := filepath.Join(home, ".talus-helper", "screenshots")
    return &ScreenshotListService{BaseDir: base}
}

var imageExts = map[string]struct{}{
    ".png":  {},
    ".jpg":  {},
    ".jpeg": {},
    ".gif":  {},
    ".webp": {},
}

// ListScreenshots returns the list of image files sorted by newest first
func (s *ScreenshotListService) ListScreenshots() ([]ScreenshotFile, error) {
    entries, err := os.ReadDir(s.BaseDir)
    if err != nil {
        if errors.Is(err, fs.ErrNotExist) {
            // Directory missing: treat as empty list
            return []ScreenshotFile{}, nil
        }
        return nil, err
    }

    files := make([]ScreenshotFile, 0, len(entries))
    for _, e := range entries {
        if e.IsDir() {
            continue
        }
        name := e.Name()
        ext := strings.ToLower(filepath.Ext(name))
        if _, ok := imageExts[ext]; !ok {
            continue
        }
        info, err := e.Info()
        if err != nil {
            continue
        }
        fullPath := filepath.Join(s.BaseDir, name)
        files = append(files, ScreenshotFile{
            Name:    name,
            Path:    fullPath,
            Size:    info.Size(),
            ModTime: info.ModTime(),
        })
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i].ModTime.After(files[j].ModTime)
    })

    return files, nil
}

// OpenScreenshot opens the given file path with the OS default application
func (s *ScreenshotListService) OpenScreenshot(path string) error {
    if path == "" {
        return errors.New("empty path")
    }
    // If relative, resolve against BaseDir
    if !filepath.IsAbs(path) {
        path = filepath.Join(s.BaseDir, path)
    }
    if _, err := os.Stat(path); err != nil {
        return err
    }

    switch runtime.GOOS {
    case "darwin":
        return exec.Command("open", path).Start()
    case "windows":
        return exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
    default:
        // linux and others
        return exec.Command("xdg-open", path).Start()
    }
}


