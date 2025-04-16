package app

import (
	"errors"
	"log"
	"math/rand"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

type SnippetService struct {
	repo domain.CodeSnippetRepository
}

func NewSnippetService(repo domain.CodeSnippetRepository) *SnippetService {
	return &SnippetService{repo: repo}
}

func (s *SnippetService) AddSnippet(repoName, fileName, language, snippetContent string) error {
	log.Printf("Storing code snippet")
	snippet := domain.CodeSnippet{
		Name:       fileName,
		Repository: repoName,
		RepoDir:    "repo_dir_example",
		Content:    snippetContent,
		Language:   language,
	}
	return s.repo.Save(snippet)
}

func (s *SnippetService) GetSnippetsByRepository(repoName string) ([]domain.CodeSnippet, error) {
	return s.repo.GetByRepository(repoName)
}

func (s *SnippetService) GetSnippetsByFileName(fileName string) ([]domain.CodeSnippet, error) {
	return s.repo.GetByFileName(fileName)
}

func (s *SnippetService) DeleteSnippet(snippetID string) error {
	return s.repo.Delete(snippetID)
}

func (s *SnippetService) GetRandomSnippetByLanguage(lang string) (*domain.CodeSnippet, error) {
	snippets, err := s.repo.GetByLanguage(lang)
	if err != nil {
		return nil, err
	}

	if len(snippets) == 0 {
		return nil, errors.New("no snippets available for this language")
	}

	index := rand.Intn(len(snippets))
	return &snippets[index], nil
}
