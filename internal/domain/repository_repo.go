package domain

type RepositoryRepo interface {
  SearchByLang(lang string) ([]Repository, error)
}
