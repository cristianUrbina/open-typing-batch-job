package repositories

import "cristianUrbina/open-typing-batch-job/internal/domain"

type RepositoryRepo interface {
  SearchByLang(lang string) ([]domain.Repository, error)
}
