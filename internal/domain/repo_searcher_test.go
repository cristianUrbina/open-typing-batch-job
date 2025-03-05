package domain

import (
	"testing"

	"cristianUrbina/open-typing-batch-job/testutils"
)

func TestSearchRepoByLang(t *testing.T) {
	// arrange
	lang := "c"
	expected, _ := CreateRepositorySlice() // act
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
	if !testutils.AreSlicesEqual(result, expected) {
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

func (d *DummyRepositoryRepo) GetRepoContent(r Repository) (*RepositoryWithContent, error) {
	return nil, nil
}

func CreateRepositorySlice() ([]Repository, error) {
	result := []Repository{
		{
			Name:   "sample",
			Author: "cristian",
			Lang:   "c",
			Source: "github",
		},
		{
			Name:   "otherrepo",
			Author: "otherauthor",
			Lang:   "c",
			Source: "github",
		},
	}
	return result, nil
}
