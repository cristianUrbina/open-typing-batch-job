package domain

import fileutils "cristianUrbina/open-typing-batch-job/pkg/filetutils"

func NewCodeFileContentReader() *CodeFileContentReader {
	return &CodeFileContentReader{}
}

type CodeFileContentReader struct{}


func (f *CodeFileContentReader) Read(repo Repository, filePath string) (*Code, error) {
  content, err := fileutils.GetFileContent(filePath)
  if err != nil {
    return nil, err
  }
  result := &Code {
    Repository: repo,
    RepoDir: "",
    Content: content,
  }
  return result, nil
}
