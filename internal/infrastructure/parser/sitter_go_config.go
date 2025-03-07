package infrastructure

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

type SitterGoConfig struct{}

func (s *SitterGoConfig) GetLanguage() *sitter.Language {
	return golang.GetLanguage()
}

func (s *SitterGoConfig) GetQuery() (*sitter.Query, error) {
	queryStr := `
	[
		(function_declaration) @func
		(method_declaration) @method
	]`
	return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterGoConfig) GetExtractor() (*Extractor, error) {
	query, err := s.GetQuery()
	if err != nil {
		return nil, err
	}
	return NewExtractor(query), nil
}
