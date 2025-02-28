package testutils

import (
	"bytes"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

func CreateRepositorySlice() ([]domain.Repository, error) {
	sampleCont, err := CreateSampleTarGZ()
	if err != nil {
	  return nil, err
	}
	result := []domain.Repository{
		{
			Name:    "cristian/sample",
			Lang:    "c",
			Source:  "github",
			Content: bytes.NewReader(sampleCont.Bytes()),
		},
		{
			Name:    "otherauthor/otherrepo",
			Lang:    "c",
			Source:  "github",
			Content: bytes.NewReader(sampleCont.Bytes()),
		},
	}
	return result, nil
}
