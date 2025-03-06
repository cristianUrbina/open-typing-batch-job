package app

import (
	"bytes"
	"testing"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"
)

type DummyCodeRepositoryRepo struct {
	CreateCalled    bool
	LastCodeProject *domain.RepositoryWithContent
	Err             error
}

func (d *DummyCodeRepositoryRepo) Create(code *domain.RepositoryWithContent) error {
	d.CreateCalled = true
	d.LastCodeProject = code
	return nil
}

func TestAddValidRepository(t *testing.T) {
	// arrange

	contents := map[string]string{
		"file1.txt": "Hello, world!",
		"file2.txt": "Go is awesome!",
	}
	tarGzData, err := testutils.CreateTarGz(contents)
	if err != nil {
		t.Fatalf("Failed to create tar.gz: %v", err)
	}
	code := &domain.RepositoryWithContent{
		Name:    "some name",
		Content: bytes.NewReader(tarGzData.Bytes()),
	}

	dummyRepo := &DummyCodeRepositoryRepo{}
	// codeProjectTarRepo :=
	svc := NewCodeProjectService(dummyRepo)

	// act
	err = svc.AddRepo(code)
	// assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !dummyRepo.CreateCalled {
		t.Error("expected Create to be called for valid code, but it wasn't")
	}
	if dummyRepo.LastCodeProject != code {
		t.Errorf("expected LastCode to be %v, got %v", code, dummyRepo.LastCodeProject)
	}
}

func TestAddInvalidRepository(t *testing.T) {
	// arrange
	code := &domain.RepositoryWithContent{
		Name: "some name",
	}

	dummyRepo := &DummyCodeRepositoryRepo{}
	// codeProjectTarRepo :=
	svc := NewCodeProjectService(dummyRepo)

	// act
	err := svc.AddRepo(code)

	if err == nil {
		t.Errorf("expected error %v, got nil", ErrInvalidCode)
	}
	if err != ErrInvalidCode {
		t.Errorf("expected error %v, got %v", ErrInvalidCode, err)
	}
	if dummyRepo.CreateCalled {
		t.Error("expected Create not to be called for invalid code, but it was")
	}
}
