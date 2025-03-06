package domain

import (
	"errors"
	"io"

	"cristianUrbina/open-typing-batch-job/pkg/ioutils"
)

type Repository struct {
	Name   string
	Author string
	Lang   string
	Source string
}

type RepositoryWithContent struct {
	Name    string
	Author  string
	Lang    string
	Source  string
	Content io.ReadSeeker
}

func (c *RepositoryWithContent) Validate() error {
	hasContent, err := ioutils.HasContent(c.Content)
	if err != nil {
		return err
	}
	if !hasContent {
		return ErrEmptyContent
	}
	return nil
}

func (r *Repository) GetFullName() string {
	return r.Name
}

func (c *Repository) Validate() error {
	return nil
}

var ErrEmptyContent = errors.New("empty content")

func (c Repository) Equal(other Repository) bool {
	if c.Name != other.Name || c.Author != other.Author || c.Lang != other.Lang || c.Source != other.Source {
		return false
	}
	return true
}
