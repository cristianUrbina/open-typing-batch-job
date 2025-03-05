package infrastructure

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

type SitterPythonConfig struct{}

func (s *SitterPythonConfig) GetLanguage() *sitter.Language {
	return python.GetLanguage()
}

func (s *SitterPythonConfig) GetQuery() (*sitter.Query, error) {
	queryStr := `
        [
            (function_definition) @func
            (class_definition) @class
        ]
    `
	return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterPythonConfig) GetExtractor() (*Extractor, error) {
	query, err := s.GetQuery()
	if err != nil {
		return nil, err
	}
	return NewExtractor(query), nil
}
