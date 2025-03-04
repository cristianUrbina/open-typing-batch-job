package domain


func NewRepoSearcher(repo RepositoryRepo) *RepoSearcher {
  return &RepoSearcher{
    repo: repo,
  }
}

type RepoSearcher struct {
  repo RepositoryRepo
}

func (uc *RepoSearcher) SearchByLang(lang string) ([]Repository, error){
  repos, err := uc.repo.SearchByLang(lang)
  if err != nil {
    return nil, err
  }
  return repos, nil
}
