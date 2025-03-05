package domain

import (
	"cristianUrbina/open-typing-batch-job/testutils"
	"testing"
)


func TestRepoFilter(t *testing.T) {
  // arrange
  files := []string{
    "/tmp/main.c",
    "/tmp/utils.c",
    "/tmp/utils.h",
    "/tmp/README.txt",
  }
  expected := []string{
    "/tmp/main.c",
    "/tmp/utils.c",
    "/tmp/utils.h",
  }
  fileFilter := NewFileFilter([]string{"c", "h"})
  // act
  filteredFiles, err := fileFilter.Filter(files)

  // assert
  if err != nil {
    t.Errorf("error was expected to be nil but got %v", err)
  }
  if !testutils.AreStrSlicesEqual(filteredFiles, expected) {
    t.Errorf("filterd files were expeted to be %v, but got %v", expected, filteredFiles)
  }
}
