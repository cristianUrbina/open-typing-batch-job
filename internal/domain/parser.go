package domain

type Parser interface {
  Parse(code []byte) ([]CodeSnippet, error)
}
