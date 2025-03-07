package app

import "cristianUrbina/open-typing-batch-job/internal/domain"

type SnippetService struct {
	repo domain.CodeSnippetRepository
}

func NewSnippetService(repo domain.CodeSnippetRepository) *SnippetService {
	return &SnippetService{repo: repo}
}

func (s *SnippetService) AddSnippet(repoName, fileName, language, snippetContent string) error {
	snippet := domain.CodeSnippet{
		Name:       fileName,
		Repository: repoName,
		RepoDir:    "repo_dir_example", // Or whatever the directory info is
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
