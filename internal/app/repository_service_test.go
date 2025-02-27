package app

import (
	"bytes"
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"
	"testing"
)

type DummyCodeRepositoryRepo struct {
  CreateCalled bool
  LastCodeProject     *domain.Repository
  Err          error
}

func (d *DummyCodeRepositoryRepo) Create(code *domain.Repository) error {
  d.CreateCalled = true
  d.LastCodeProject = code
  return nil
}

func TestAddValidRepository(t *testing.T) {
  // arrange
  tarGZ, tarErr := testutils.CreateSampleTarGZ()
  if tarErr != nil { t.Fatal(tarErr) }
  code := &domain.Repository{
    Name: "some name",
    Content: bytes.NewReader(tarGZ.Bytes()),
  }

  dummyRepo := &DummyCodeRepositoryRepo{}
  // codeProjectTarRepo :=
  svc := NewCodeProjectService(dummyRepo)

  // act
  err := svc.AddRepo(code)

  //assert
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
  code := &domain.Repository{
    Name: "some name",
    Content: bytes.NewReader([]byte{}),
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
