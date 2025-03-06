package githubrepo

import (
	"fmt"
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
		repository := domain.Repository{
			Name:   v.FullName,
			Lang:   lang,
			Source: "github",
		}
		repos = append(repos, repository)
	}
	return repos, nil
}

func (r *RepositoryGitHHubRepo) GetRepoContent(repo domain.Repository) (*domain.RepositoryWithContent, error) {
	tarballFile, err := r.client.GetRepoTarball(repo.GetFullName())
	if err != nil {
		return nil, fmt.Errorf("error getting repo tarball content: %w", err)
	}
	// var buf bytes.Buffer
	// n, err := io.CopyN(&buf, tarballFile, 512) // Read first 512 bytes
	// if err != nil && err != io.EOF {
	// 	log.Fatal("Error reading tarball:", err)
	// 	return nil, nil
	// }
	// if n == 0 {
	// 	log.Fatal("Tarball is empty")
	// 	return nil, nil
	// }
	// log.Printf("tarball length %v", n)
	return &domain.RepositoryWithContent{
		Name:    repo.Name,
		Author:  repo.Author,
		Lang:    repo.Lang,
		Source:  repo.Source,
		Content: tarballFile,
	}, nil
}
