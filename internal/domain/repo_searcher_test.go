package domain

import (
	"bytes"
	"testing"

	"cristianUrbina/open-typing-batch-job/testutils"
)

func TestSearchRepoByLang(t *testing.T) {
	// arrange
	lang := "c"
  expected, _ := CreateRepositorySlice()  // act
	dummyRepositoryRepo := &DummyRepositoryRepo{
		Repos: expected,
	}
	useCase := NewRepoSearcher(dummyRepositoryRepo)

	// act
	result, err := useCase.SearchByLang(lang)

	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if result == nil {
		t.Errorf("result was expected to not be nil")
	}
	if !areRepoSlicesEqual(result, expected) {
		t.Errorf("result expected %v, but got %v", expected, result)
	}
}

type DummyRepositoryRepo struct {
	LastLang string
	Repos    []Repository
}

func (d *DummyRepositoryRepo) SearchByLang(lang string) ([]Repository, error) {
	d.LastLang = lang
	return d.Repos, nil
}

func areRepoSlicesEqual(a, b []Repository) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equal(&b[i]) {
			return false
		}
	}

	return true
}

func CreateRepositorySlice() ([]Repository, error) {
	sampleCont, err := testutils.CreateSampleTarGZ()
	if err != nil {
	  return nil, err
	}
	result := []Repository{
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
