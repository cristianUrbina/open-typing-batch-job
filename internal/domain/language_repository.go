package domain

type LanguageRepository interface {
	GetLanguages() ([]Language, error)
	GetLanguageByAlias(name string)(*Language, error)
}
