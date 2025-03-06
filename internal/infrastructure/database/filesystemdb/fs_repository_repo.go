package filesystemdb

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	fileutils "cristianUrbina/open-typing-batch-job/pkg/filetutils"
	"os"
	"path/filepath"
	"strings"
)

type FSRepositoryRepo struct { }

func (f *FSRepositoryRepo) Create(code *domain.RepositoryWithContent) error {
  tmpDir := os.TempDir()
  repoDir := strings.Split(code.Name, "/")
  path := filepath.Join(tmpDir, "repos", code.Source, code.Lang, repoDir[0], repoDir[1]+".tar.gz")
  fileutils.SaveContentToFile(code.Content, path)
  return nil
}
