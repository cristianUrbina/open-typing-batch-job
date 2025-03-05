package domain

import (
	"cristianUrbina/open-typing-batch-job/pkg/filetutils"
	"fmt"
)

func NewCodeExtractor(dir string) *TarballExtractor {
  return &TarballExtractor{
  	dir: dir,
  }
}

type TarballExtractor struct {
	dir string
}

func (c *TarballExtractor) Extract(r *Repository) ([]string, error) {
  files, err := fileutils.ExtractTarball(r.Content, c.dir)
  if err != nil {
	  return nil, fmt.Errorf("error extracting temporary directory %v", err)
  }
  return files, nil
}
