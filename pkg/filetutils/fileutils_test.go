package fileutils

import (
	"bytes"
	"cristianUrbina/open-typing-batch-job/testutils"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractTar(t *testing.T) {
	// arrange
	tempDir, err := os.MkdirTemp("", "test-extract")
	if err != nil {
		t.Fatalf("Failed to created tempd dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	contents := map[string]string{
		"file1.txt": "Hello, world!",
		"file2.txt": "Go is awesome!",
	}
	tarGzData, err := testutils.CreateTarGz(contents)
	if err != nil {
		t.Fatalf("Failed to create tar.gz: %v", err)
	}

	// act
	err = ExtractTarball(bytes.NewReader(tarGzData.Bytes()), tempDir)
	if err != nil {
		t.Fatalf("ExtractTarball failed: %v", err)
	}

	// assert
	for name, expectedContent := range contents {
		path := filepath.Join(tempDir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("Failed to read extracted file: %v", err)
		}

		if string(data) != expectedContent {
			t.Errorf("File content mismatch: expected %q, got %q", expectedContent, string(data))
		}
	}
}
