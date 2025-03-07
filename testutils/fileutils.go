package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateTempFiles(files map[string]string) (string, func(), error) {
	tempDir, err := os.MkdirTemp("", "testfiles")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	for name, content := range files {
		filePath := filepath.Join(tempDir, name)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return "", nil, fmt.Errorf("failed to create file %s: %w", name, err)
		}
	}

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup, nil
}
