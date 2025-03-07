package domain

type LanguageRepository interface {
	GetLanguages() ([]Language, error)
}
