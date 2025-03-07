package app

import (
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

func NewCodeService(s domain.SnippetExtractor) *CodeService {
	return &CodeService{
		snippetExtractor: s,
	}
}

type CodeService struct {
	snippetExtractor domain.SnippetExtractor
}

func (c *CodeService) Analyze(code *domain.Code) ([]domain.CodeSnippet, error) {
	snippets, err := c.snippetExtractor.ExtractSnippets(code)
	if err != nil {
		return nil, fmt.Errorf("error extracting snippets: %w", err)
	}
	var codeSnippets []domain.CodeSnippet
	for _, s := range snippets {
		codeSnippets = append(codeSnippets, domain.CodeSnippet{
			Content:    s.Content,
			Language:   code.Repository.Lang.Alias,
			Repository: code.Repository.GetFullName(),
			RepoDir:    code.RepoDir,
		},
		)
	}
	return codeSnippets, nil
}
