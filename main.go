package main

import (
	"log"

	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
	"cristianUrbina/open-typing-batch-job/pkg/filetutils"
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
			codeDir := "/tmp/repos/"+l
			// code := &entitites.Code{
			// 	Name: v.FullName,
			// 	Content: tarballFile,
			// }
			log.Printf("Extracting tarbal for repo: %+v\n", v.FullName)
			fileutils.ExtractTarball(tarballFile, codeDir)
		}
	}
}
