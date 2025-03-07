package app

import "cristianUrbina/open-typing-batch-job/internal/domain"

type LanguageService struct {
	Repo domain.LanguageRepository
}

func (s *LanguageService) GetAvailableLanguages() ([]domain.Language, error) {
	return s.Repo.GetLanguages()
}
