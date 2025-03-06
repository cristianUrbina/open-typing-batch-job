package domain

type SnippetExtractor interface {
	ExtractSnippets(code *Code) ([]Snippet, error)
}
