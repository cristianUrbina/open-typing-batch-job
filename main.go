package main

import (
	"log"
	"os"

	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/filesystemdb"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/githubrepo"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

func main() {
	log.Printf("Searching for github repos")
	languages := []string{"c", "python", "go", "js", "rust", "java"}
	for _, l := range languages {
		client := githubapiclient.NewAPIClient()
		githubRepo := githubrepo.NewRepositoryGithubRepo(*client)
		repoSearcher := domain.NewRepoSearcher(githubRepo)
		tmpDir, err := os.MkdirTemp("", "test-extract")
		if err != nil {
			log.Fatalf("failed to create temporary directory")
		}
		repoExtractor := domain.NewCodeExtractor(tmpDir)
		repos, err := repoSearcher.SearchByLang(l)
		if err != nil {
			log.Fatalf("Failed searching repos %v", err)
		}
		for _, r := range repos {
			fsRepo := &filesystemdb.FSRepositoryRepo{}
			svc := app.NewCodeProjectService(fsRepo)
			err = svc.AddRepo(&r)
			if err != nil {
				log.Fatalf("an unexpected error ocurred while adding code, %v", err)
			}
			log.Printf("extracting repo %v tarball", r.Name)
			_, err := repoExtractor.Extract(&r)
			if err != nil {
				log.Fatalf("error while extracting the tarball %v", err)
			}
		}
	}
}
