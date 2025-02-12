package fileutils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

func createTarGz(contents map[string]string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	gzWriter := gzip.NewWriter(buf)
	tarWriter := tar.NewWriter(gzWriter)

	for name, content := range contents {
		header := &tar.Header{
			Name: name,
			Size: int64(len(content)),
			Mode: 0o644,
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := tarWriter.Write([]byte(content)); err != nil {
			return nil, err
		}
	}

	tarWriter.Close()
	gzWriter.Close()

	return buf, nil
}

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
	tarGzData, err := createTarGz(contents)
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
