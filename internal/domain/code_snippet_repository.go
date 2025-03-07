package domain

type CodeSnippetRepository interface {
	Save(snippet CodeSnippet) error
	GetByRepository(repoName string) ([]CodeSnippet, error)
	GetByFileName(fileName string) ([]CodeSnippet, error)
	Delete(snippetID string) error
}
