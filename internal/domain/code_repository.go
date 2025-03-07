package domain

type CodeRepository interface {
  Create(CodeProject *RepositoryWithContent) error
}
