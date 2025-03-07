package infrastructure

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/rust"
)

type SitterRustConfig struct{}

func (s *SitterRustConfig) GetLanguage() *sitter.Language {
	return rust.GetLanguage()
}

func (s *SitterRustConfig) GetQuery() (*sitter.Query, error) {
	queryStr := `[
		(function_item) @func
		(struct_item) @struct
	]`
	return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterRustConfig) GetExtractor() (*Extractor, error) {
	query, err := s.GetQuery()
	if err != nil {
		return nil, err
	}
	return NewExtractor(query), nil
}
