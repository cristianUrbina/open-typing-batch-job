package app

import (
	"fmt"
	"os"

	"cristianUrbina/open-typing-batch-job/internal/domain"

	fileutils "cristianUrbina/open-typing-batch-job/pkg/filetutils"
)

func NewRepoService(repo domain.RepositoryRepo) *RepoService {
	return &RepoService{
		repo: repo,
	}
}

type RepoService struct {
	repo domain.RepositoryRepo
}

func (r *RepoService) SearchByLang(lang *domain.Language) ([]domain.Repository, error) {
	repos, err := r.repo.SearchByLang(lang)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (r *RepoService) GetRepoContent(repository domain.Repository) (*domain.RepositoryWithContent, error) {
	return r.repo.GetRepoContent(repository)
}

func (c *RepoService) Extract(r *domain.RepositoryWithContent) ([]string, error) {
	tmpDir, err := os.MkdirTemp("", "test-extract")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary dir: %v", err)
	}
	return fileutils.ExtractTarball(r.Content, tmpDir)
}
