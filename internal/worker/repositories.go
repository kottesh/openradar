package worker

// This handles Repository cloning
// for the worker package

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"openradar/internal/config"
)

// Directories to skip during scanning
var skipDirectories = map[string]struct{}{
	".git":         {},
	"node_modules": {},
	"vendor":       {},
	"dist":         {},
	"build":        {},
	".terraform":   {},
}

// File extensions to scan
var toScan = map[string]struct{}{
	".env":  {},
	".md":   {},
	".txt":  {},
	".py":   {},
	".rs":   {},
	".yml":  {},
	".ts":   {},
	".js":   {},
	".yaml": {},
	".go":   {},
	".json": {},
	".toml": {},
	".php":  {},
	".rb":   {},
	".java": {},
	".kt":   {},
	".sh":   {},
}

type ScannedFile struct {
	RelPath string
	Content string
}

// Scan files in repository and run callback per file
func ScanFiles(directory string, cfg config.Config) ([]ScannedFile, error) {

	// Func here
	var buf bytes.Buffer
	var results []ScannedFile
	maxSize := int64(cfg.Scanner.MaxFileSizeKB * 1024)

	// Crawl directories
	err := filepath.WalkDir(directory, func(path string, dir os.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		// Check if is directory & should be skipped.
		if dir.IsDir() {

			if _, ok := skipDirectories[dir.Name()]; ok {

				return filepath.SkipDir
			}

			return nil
		}

		// Never follow symlinks from an untrusted repository. A repository could
		// otherwise point a scannable filename at a file outside its clone.
		if dir.Type()&os.ModeSymlink != 0 {
			return nil
		}

		// File has target extension
		ext := strings.ToLower(filepath.Ext(dir.Name()))
		_, ok := toScan[ext]
		if !ok {

			return nil
		}

		// Information exists
		inf, err := dir.Info()
		if err != nil {

			return nil
		}

		// Check size
		if inf.Size() > maxSize {

			return nil
		}

		open, err := os.Open(path)
		if err != nil {

			return nil
		}

		buf.Reset()

		_, err = io.Copy(&buf, open)

		open.Close()

		if err != nil {

			return nil
		}

		relPath, _ := filepath.Rel(directory, path)
		results = append(results, ScannedFile{RelPath: relPath, Content: buf.String()})
		return nil
	})

	return results, err
}

// Clone repository
func CloneRepo(Context context.Context, cloneURL string, dir string) error {

	// Check if URL is malicious/non-http compliant
	if !strings.HasPrefix(cloneURL, "https://") {

		return fmt.Errorf("Non HTTP compliant url! %s", cloneURL)
	}

	// Execute command to clone
	cmd := exec.CommandContext(Context, "git", "clone", "--depth", "1", "--single-branch", cloneURL, dir)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
