package domain

import "fmt"

func NewCodeAnalyzer(s SnippetExtractor) *CodeAnalyzer {
	return &CodeAnalyzer{
		snippetExtractor: s,
	}
}

type CodeAnalyzer struct {
	snippetExtractor SnippetExtractor
}

func (c *CodeAnalyzer) Analyze(code *Code) ([]CodeSnippet, error) {
	snippets, err := c.snippetExtractor.ExtractSnippets(code)
	if err != nil {
		return nil, fmt.Errorf("error extracting snippets: %w", err)
	}
	var codeSnippets []CodeSnippet
	for _, s := range snippets {
		codeSnippets = append(codeSnippets, CodeSnippet{
			Content:    s.Content,
			Language:   code.Repository.Lang,
			Repository: code.Repository.GetFullName(),
			RepoDir:    code.RepoDir,
		},
		)
	}
	return codeSnippets, nil
}
