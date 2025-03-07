package infrastructure

import (
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"

	sitter "github.com/smacker/go-tree-sitter"
)

type SnippetExtractor interface {
	Extract(code *domain.Code, rootNode *sitter.Node) ([]domain.Snippet, error)
}

type SitterLangConfig interface {
	GetLanguage() *sitter.Language
	GetQuery() (*sitter.Query, error)
	GetExtractor() (*Extractor, error)
}

func NewLangSitterConfig(lang string) (SitterLangConfig, error) {
	switch lang {
	case "python":
		return &SitterPythonConfig{}, nil
	case "java":
		return &SitterJavaConfig{}, nil
	case "javascript":
		return &SitterJavascriptConfig{}, nil
	default:
		return nil, fmt.Errorf("unsupported language %v", lang)
	}
}

type ExtractorLanguageConfig struct {
	language *sitter.Language
	query    string
}
