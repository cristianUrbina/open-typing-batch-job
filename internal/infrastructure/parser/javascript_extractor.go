package infrastructure

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"fmt"

	"github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

type JavaScriptFunctionExtractor struct{}

func NewJavaScriptFunctionExtractor() *JavaScriptFunctionExtractor {
    return &JavaScriptFunctionExtractor{}
}

func (j *JavaScriptFunctionExtractor) Extract(code *domain.Code, rootNode *sitter.Node) ([]domain.Snippet, error) {
    query, err := sitter.NewQuery([]byte("(function_declaration) @func"), javascript.GetLanguage())
    if err != nil {
        return nil, fmt.Errorf("failed to create query: %w", err)
    }
    defer query.Close()

    cursor := sitter.NewQueryCursor()
    defer cursor.Close()

    cursor.Exec(query, rootNode)

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
