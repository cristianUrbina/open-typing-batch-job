package main

import (
	"log"

	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/filesystemdb"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

func main() {
	log.Printf("Searching for github repos")
	languages := []string{"c", "python", "go", "js", "rust", "java"}
	for _, l := range languages {
		apiClient := githubapiclient.NewAPIClient()
		searchResp, err := githubapiclient.SearchGitHubRepos(apiClient, l)
		if err != nil {
			log.Fatalf("Failed to search github repos: %v", err)
		}
		for _, v := range searchResp.Items {
			log.Printf("Getting tarbal for repo: %+v\n", v.FullName)
			tarballFile, err := githubapiclient.GetRepoTarball(v.FullName)
			if err != nil {
				return
			}
			codeProj := &domain.Repository{
				Name:    v.FullName,
				Lang:    l,
				Source:  "github",
				Content: tarballFile,
			}
			repo := &filesystemdb.FSRepositoryRepo{}
			svc := app.NewCodeProjectService(repo)
			err = svc.AddRepo(codeProj)
			if err != nil {
				log.Fatalf("an unexpected error ocurred while adding code, %v", err)
			}
		}
	}
}
