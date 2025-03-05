package infrastructure

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"

	sitter "github.com/smacker/go-tree-sitter"
)

func NewExtractor(query *sitter.Query) *Extractor {
	return &Extractor{
		query: query,
	}
}

type Extractor struct {
	query *sitter.Query
}

func (f *Extractor) Extract(code *domain.Code, rootNode *sitter.Node) ([]domain.Snippet, error) {
	defer f.query.Close()

	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	cursor.Exec(f.query, rootNode)

	var functions []domain.Snippet
	for {
		match, ok := cursor.NextMatch()
		if !ok {
			break
		}

		for _, capture := range match.Captures {
			functionContent := capture.Node.Content(code.Content)
			functions = append(functions, domain.Snippet{
				Content: functionContent,
			})
		}
	}
	return functions, nil
}
