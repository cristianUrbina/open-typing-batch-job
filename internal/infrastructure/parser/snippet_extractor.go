package infrastructure

import (
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/python"
)

type SnippetExtractor interface {
	Extract(code *domain.Code, rootNode *sitter.Node) ([]domain.Snippet, error)
}

// type SitterByLang struct {
// 	Language  *sitter.Language
// 	Extractor SnippetExtractor
// }

type SitterLangConfig interface {
	GetLanguage() *sitter.Language
	GetQuery() (*sitter.Query, error)
	GetExtractor() (*Extractor, error) // Returns *Extractor
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

func getLangCapabilities(lang string) (*ExtractorLanguageConfig, error) {
	var query string
	var language *sitter.Language
	switch lang {
	case "python":
		language = python.GetLanguage()
		query = `
            [
                (function_definition) @func
                (class_definition) @class
            ]
        `
	case "java":
		language = java.GetLanguage()
		query = `[
                (method_declaration) @method
            ]`
	case "javascript":
		language = javascript.GetLanguage()
		query = `[
                (function_declaration) @func
                (function_expression) @func
                (method_definition) @method
            ]`
	default:
		language = python.GetLanguage()
	}

	return &ExtractorLanguageConfig{
		language: language,
		query:    query,
	}, nil
}

type ExtractorLanguageConfig struct {
	language *sitter.Language
	query    string
}
