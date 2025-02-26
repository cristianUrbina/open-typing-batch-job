package services

import (
	"bytes"
	"cristianUrbina/open-typing-batch-job/internal/core/entities"
	"cristianUrbina/open-typing-batch-job/testutils"
	"testing"
)

type DummyCodeRepo struct {
  CreateCalled bool
  LastCode     *entitites.Code
  Err          error
}

func (d *DummyCodeRepo) Create(code *entitites.Code) error {
  d.CreateCalled = true
  d.LastCode = code
  return nil
}

func TestAddValidCode(t *testing.T) {
  // arrange
  tarGZ, tarErr := testutils.CreateSampleTarGZ()
  if tarErr != nil {
    t.Fatal(tarErr)
  }

  code := &entitites.Code{
    Name: "some name",
    Content: bytes.NewReader(tarGZ.Bytes()),
  }

  dummyRepo := &DummyCodeRepo{}
  // codeProjectTarRepo :=
  svc := NewCodeService(dummyRepo)

  // act
  err := svc.AddCode(code)

  //assert
  if err != nil {
    t.Errorf("expected no error, got %v", err)
  }

  if !dummyRepo.CreateCalled {
    t.Error("expected Create to be called for valid code, but it wasn't")
  }
  if dummyRepo.LastCode != code {
    t.Errorf("expected LastCode to be %v, got %v", code, dummyRepo.LastCode)
  }
}

func TestAddInvalidCode(t *testing.T) {
  // arrange
  code := &entitites.Code{
    Name: "some name",
    Content: bytes.NewReader([]byte{}),
  }

  dummyRepo := &DummyCodeRepo{}
  // codeProjectTarRepo :=
  svc := NewCodeService(dummyRepo)

  // act
  err := svc.AddCode(code)

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
