package infrastructure

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

type SitterJavascriptConfig struct{}

func (s *SitterJavascriptConfig) GetLanguage() *sitter.Language {
	return javascript.GetLanguage()
}

func (s *SitterJavascriptConfig) GetQuery() (*sitter.Query, error) {
	queryStr := `[
    (function_declaration) @func
    (function_expression) @func
    (method_definition) @method
  ]`
	return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterJavascriptConfig) GetExtractor() (*Extractor, error) {
	query, err := s.GetQuery()
	if err != nil {
		return nil, err
	}
	return NewExtractor(query), nil
}
