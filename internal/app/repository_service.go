package app

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/domain/repositories"
	"errors"
)

func NewCodeProjectService(repo repositories.CodeRepository) *RepositoryService{
  return &RepositoryService{
    repo: repo,
  }
}

type RepositoryService struct {
  repo repositories.CodeRepository
}

func (c *RepositoryService) AddRepo(code *domain.Repository) error {
  if err := code.Validate(); err != nil {
    return ErrInvalidCode
  }
  return c.repo.Create(code)
}

var ErrInvalidCode = errors.New("invalid code data")
