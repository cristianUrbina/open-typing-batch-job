package domain

type LanguageService struct {
	Repo LanguageRepository
}

func (s *LanguageService) GetAvailableLanguages() ([]Language, error) {
	return s.Repo.GetLanguages()
}
