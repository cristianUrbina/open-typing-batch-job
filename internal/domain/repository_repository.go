package domain

type RepositoryRepo interface {
  SearchByLang(lang *Language) ([]Repository, error)
  GetRepoContent(r Repository) (*RepositoryWithContent, error)
}
