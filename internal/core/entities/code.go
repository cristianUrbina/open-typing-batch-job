package entitites

import (
	"cristianUrbina/open-typing-batch-job/pkg/ioutils"
	"errors"
	"io"
	// "cristianUrbina/open-typing-batch-job/pkg/ioutils"
)

type Code struct {
  Name string
  Content io.ReadSeeker
}

func (c *Code) Validate() error {
  hasContent, err := ioutils.HasContent(c.Content)
  if err != nil {
    return err
  }
  if !hasContent {
    return ErrEmptyContent
  }
  return nil
}

var ErrEmptyContent = errors.New("empty content")

