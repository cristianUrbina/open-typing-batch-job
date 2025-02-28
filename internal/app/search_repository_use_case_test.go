package app

import (
	"testing"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"
)

func TestSearchRepoByLang(t *testing.T) {
	// arrange
	lang := "c"
  expected, err := testutils.CreateRepositorySlice()  // act
	dummyRepositoryRepo := &DummyRepositoryRepo{
		Repos: expected,
	}
	useCase := NewSearchReposByLangUseCase(dummyRepositoryRepo)

	// act
	result, err := useCase.Execute(lang)

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
	Repos    []domain.Repository
}

func (d *DummyRepositoryRepo) SearchByLang(lang string) ([]domain.Repository, error) {
	d.LastLang = lang
	return d.Repos, nil
}

func areRepoSlicesEqual(a, b []domain.Repository) bool {
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
