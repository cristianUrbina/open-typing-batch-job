package domain

type CodeRepository interface {
  Create(CodeProject *Repository) error
}
