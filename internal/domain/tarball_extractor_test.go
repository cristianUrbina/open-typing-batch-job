package domain

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"cristianUrbina/open-typing-batch-job/testutils"
)

func TestCodeExtractorExtract(t *testing.T) {
	// arrange
	contents := map[string]string{
		"file1.txt": "Hello, world!",
		"file2.txt": "Go is awesome!",
	}
	tmpDir, err := os.MkdirTemp("", "test-extract")
	if err != nil {
		t.Fatalf("Failed to create temporary dir %v", err)
	}
	tarGzData, err := testutils.CreateTarGz(contents)
	if err != nil {
		t.Fatalf("Failed to create tar.gz: %v", err)
	}
	repo := &RepositoryWithContent{
		Name:    "repo",
		Author:  "anauthor",
		Lang:    "go",
		Source:  "github",
		Content: bytes.NewReader(tarGzData.Bytes()),
	}
	expectedFiles := []string{filepath.Join(tmpDir, "file1.txt"),filepath.Join(tmpDir, "file2.txt")}
	codeExtractor := NewCodeExtractor(tmpDir)
	// act
	files, err := codeExtractor.Extract(repo)

	// assert
	if err != nil {
		t.Errorf("Error was expected to be nil, but got %v", err)
	}
	if !testutils.AreStrSlicesEqual(files, expectedFiles) {
		t.Errorf("result was expected to be %v, but got %v", expectedFiles, files)
	}
}
