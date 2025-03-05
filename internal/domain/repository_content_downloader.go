package domain

func NewRepositoryContentDownloader(repo RepositoryRepo) *RepositoryContentDownloader {
  return &RepositoryContentDownloader{
    repo: repo,
  }
}

type RepositoryContentDownloader struct {
	repo RepositoryRepo
}

func (r *RepositoryContentDownloader) GetRepoContent(repository Repository) (*RepositoryWithContent, error) {
	return r.repo.GetRepoContent(repository)
}
