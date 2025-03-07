package infrastructure

import (
	"context"
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"

	sitter "github.com/smacker/go-tree-sitter"
)

func NewTreeSitterSnippetExtractor() *TreeSitterSnippetExtractor {
	return &TreeSitterSnippetExtractor{}
}

type TreeSitterSnippetExtractor struct{}

func (t *TreeSitterSnippetExtractor) ExtractSnippets(code *domain.Code) ([]domain.Snippet, error) {
	parser := sitter.NewParser()
	defer parser.Close()
	sitterConfig, err := NewLangSitterConfig(code.Repository.Lang.Alias)
	if err != nil {
		return nil, fmt.Errorf("error creating sitteraux: %w", err)
	}
	parser.SetLanguage(sitterConfig.GetLanguage())
	ctx := context.Background()
	tree, err := parser.ParseCtx(ctx, nil, code.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse code: %w", err)
	}
	rootNode := tree.RootNode()
	extr, err := sitterConfig.GetExtractor()
	if err != nil {
		return nil, fmt.Errorf("error creating the extractor: %w", err)
	}
	return extr.Extract(code, rootNode)
}
