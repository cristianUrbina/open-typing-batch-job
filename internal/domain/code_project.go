package domain

import (
	"bytes"
	"cristianUrbina/open-typing-batch-job/pkg/ioutils"
	"errors"
	"io"
)

type Repository struct {
  Name string
  Lang string
  Source string
  Content io.ReadSeeker
}

func (c *Repository) Validate() error {
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

func (c *Repository) Equal(other *Repository) bool {
	if c == other {
		return true
	}

	if c.Name != other.Name || c.Lang != other.Lang || c.Source != other.Source {
		return false
	}

	cContent, err1 := io.ReadAll(c.Content)
	otherContent, err2 := io.ReadAll(other.Content)

	if err1 != nil || err2 != nil {
		return false
	}

	c.Content.Seek(0, io.SeekStart)
	other.Content.Seek(0, io.SeekStart)

	return bytes.Equal(cContent, otherContent)
}
