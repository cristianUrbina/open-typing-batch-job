package app

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/domain/repositories"
)

func NewSearchReposByLangUseCase(repo repositories.RepositoryRepo) *SearchReposByLangUseCase {
  return &SearchReposByLangUseCase{
    repo: repo,
  }
}

type SearchReposByLangUseCase struct {
  repo repositories.RepositoryRepo
}

func (uc *SearchReposByLangUseCase) Execute(lang string) ([]domain.Repository, error){
  repos, err := uc.repo.SearchByLang(lang)
  if err != nil {
    return nil, err
  }
  return repos, nil
}
