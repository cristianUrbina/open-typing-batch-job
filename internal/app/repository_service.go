package app

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"errors"
)

func NewCodeProjectService(repo domain.CodeRepository) *RepositoryService{
  return &RepositoryService{
    repo: repo,
  }
}

type RepositoryService struct {
  repo domain.CodeRepository
}

func (c *RepositoryService) AddRepo(code *domain.Repository) error {
  if err := code.Validate(); err != nil {
    return ErrInvalidCode
  }
  return c.repo.Create(code)
}

var ErrInvalidCode = errors.New("invalid code data")
