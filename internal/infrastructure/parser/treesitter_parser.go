package infrastructure

import (
	"context"
	"cristianUrbina/open-typing-batch-job/internal/domain"

	"github.com/smacker/go-tree-sitter"
)

type TreeSitterParser struct {

}

/*
Code {
  lang
  repo
  content
}
*/

func (t *TreeSitterParser) Parse(code domain.Code) ([]domain.CodeSnippet, error) {
  parser := sitter.NewParser()
  ctx := context.Background()
  // tree, err := parser.ParseCtx(ctx, nil, code)
  _, err := parser.ParseCtx(ctx, nil, code.Content)
  if err != nil {
    return nil, err
  }
  // rootNode := tree.RootNode()
  return nil, nil
}
