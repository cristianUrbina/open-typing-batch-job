package services

import (
	"cristianUrbina/open-typing-batch-job/internal/core/entities"
	"cristianUrbina/open-typing-batch-job/internal/core/repositories"
	"errors"
)

func NewCodeService(repo repositories.CodeRepository) *CodeService{
  return &CodeService{
    repo: repo,
  }
}

type CodeService struct {
  repo repositories.CodeRepository
}

func (c *CodeService) AddCode(code *entitites.Code) error {
  if err := code.Validate(); err != nil {
    return ErrInvalidCode
  }
  return c.repo.Create(code)
}

var ErrInvalidCode = errors.New("invalid code data")
