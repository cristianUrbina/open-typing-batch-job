package infrastructure

import (
    sitter "github.com/smacker/go-tree-sitter"
    "github.com/smacker/go-tree-sitter/java"
)

type SitterJavaConfig struct{}

func (s *SitterJavaConfig) GetLanguage() *sitter.Language {
    return java.GetLanguage()
}

func (s *SitterJavaConfig) GetQuery() (*sitter.Query, error) {
    queryStr := `
        [
            (method_declaration) @method
        ]
    `
    return sitter.NewQuery([]byte(queryStr), s.GetLanguage())
}

func (s *SitterJavaConfig) GetExtractor() (*Extractor, error) {
    query, err := s.GetQuery()
    if err != nil {
        return nil, err
    }
    return NewExtractor(query), nil
}
