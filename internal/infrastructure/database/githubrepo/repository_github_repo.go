package githubrepo

import (
	"log"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
)

func NewRepositoryGithubRepo(apiClient githubapiclient.APIClient) *RepositoryGitHHubRepo {
	return &RepositoryGitHHubRepo{
	  client: apiClient,
	}
}

type RepositoryGitHHubRepo struct {
	client githubapiclient.APIClient
}

func (r *RepositoryGitHHubRepo) SearchByLang(lang string) ([]domain.Repository, error) {
	searchResp, err := r.client.SearchGitHubRepos(lang)
	if err != nil {
		log.Fatalf("Failed to search github repos: %v", err)
	}
	repos := []domain.Repository{}
	for _, v := range searchResp.Items {
			tarballFile, err := r.client.GetRepoTarball(v.FullName)
			if err != nil {
				return nil, err
			}
			repository := domain.Repository{
				Name:    v.FullName,
				Lang:    lang,
				Source:  "github",
				Content: tarballFile,
			}
	  repos = append(repos, repository)
	}
	return repos, nil
}
