package app

import (
	"bytes"
	"path/filepath"
	"testing"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"
)

func TestSearchRepoByLang(t *testing.T) {
	// arrange
	lang := MakeLangC()
	expected, _ := CreateRepositorySlice() // act
	dummyRepositoryRepo := &DummyRepositoryRepo{
		Repos: expected,
	}
	useCase := NewRepoService(dummyRepositoryRepo)

	// act
	result, err := useCase.SearchByLang(lang)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if result == nil {
		t.Errorf("result was expected to not be nil")
	}
	if !testutils.AreSlicesEqual(result, expected) {
		t.Errorf("result expected %v, but got %v", expected, result)
	}
}

type DummyRepositoryRepo struct {
	LastLang *domain.Language
	Repos    []domain.Repository
}

func (d *DummyRepositoryRepo) SearchByLang(lang *domain.Language) ([]domain.Repository, error) {
	d.LastLang = lang
	return d.Repos, nil
}

func (d *DummyRepositoryRepo) GetRepoContent(r domain.Repository) (*domain.RepositoryWithContent, error) {
	return nil, nil
}

func MakeLangC() *domain.Language {
	return &domain.Language{
		Name: "C",
		Alias: "c",
	}
}

func CreateRepositorySlice() ([]domain.Repository, error) {
	result := []domain.Repository{
		{
			Name:   "sample",
			Author: "cristian",
			Lang:   MakeLangC(),
			Source: "github",
		},
		{
			Name:   "otherrepo",
			Author: "otherauthor",
			Lang:   MakeLangC(),
			Source: "github",
		},
	}
	return result, nil
}

func TestRepoServiceExtract(t *testing.T) {
	// arrange
	contents := map[string]string{
		"file1.txt": "Hello, world!",
		"file2.txt": "Go is awesome!",
	}
	tarGzData, err := testutils.CreateTarGz(contents)
	if err != nil {
		t.Fatalf("Failed to create tar.gz: %v", err)
	}
	repo := &domain.RepositoryWithContent{
		Name:    "repo",
		Author:  "anauthor",
		Lang:    MakeLangC(),
		Source:  "github",
		Content: bytes.NewReader(tarGzData.Bytes()),
	}
	expectedFiles := []string{"file1.txt", "file2.txt"}
	dummyRepo := &DummyRepositoryRepo{}
	svc := NewRepoService(dummyRepo)

	// act
	files, err := svc.Extract(repo)
	// assert
	if err != nil {
		t.Errorf("Error was expected to be nil, but got %v", err)
	}

	var fileNames []string
	for _, filePath := range files {
		fileNames = append(fileNames, filepath.Base(filePath))
	}
	if !testutils.AreStrSlicesEqual(fileNames, expectedFiles) {
		t.Errorf("result was expected to be %v, but got %v", expectedFiles, files)
	}
}
