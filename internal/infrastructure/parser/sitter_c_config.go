package infrastructure

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/c"
)

type SitterCConfig struct{}

func (s *SitterCConfig) GetLanguage() *sitter.Language {
	return c.GetLanguage()
}

func (s *SitterCConfig) GetQuery() (*sitter.Query, error) {
	queryStr := `
	[
		(function_definition) @func
	]`
	return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterCConfig) GetExtractor() (*Extractor, error) {
	query, err := s.GetQuery()
	if err != nil {
		return nil, err
	}
	return NewExtractor(query), nil
}
