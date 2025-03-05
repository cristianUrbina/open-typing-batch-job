package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateTempFiles(files map[string]string) (string, func(), error) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testfiles")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	// Create files inside the temp directory
	for name, content := range files {
		filePath := filepath.Join(tempDir, name)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return "", nil, fmt.Errorf("failed to create file %s: %w", name, err)
		}
	}

	// Cleanup function to remove temp directory
	cleanup := func() {
		os.RemoveAll(tempDir) // Ensure the temp directory is deleted after tests
	}

	return tempDir, cleanup, nil
}
