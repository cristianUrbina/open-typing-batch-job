package domain

import (
	"cristianUrbina/open-typing-batch-job/pkg/filetutils"
	"fmt"
)

func NewCodeExtractor(dir string) *CodeExtractor {
  return &CodeExtractor{
  	dir: dir,
  }
}

type CodeExtractor struct {
	dir string
}

func (c *CodeExtractor) Extract(r *Repository) ([]string, error) {
  files, err := fileutils.ExtractTarball(r.Content, c.dir)
  if err != nil {
	  return nil, fmt.Errorf("error extracting temporary directory %v", err)
  }
  return files, nil
}
