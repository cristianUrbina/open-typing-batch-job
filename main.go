package main

import (
	"log"

	"cristianUrbina/open-typing-batch-job/fileutils"
	"cristianUrbina/open-typing-batch-job/githubapi"
)

func main() {
	log.Printf("Searching for github repos")
	languages := []string{"c", "python", "go", "js", "rust", "java"}
	for _, l := range languages {
		apiClient := githubapi.NewAPIClient()
		searchResp, err := githubapi.SearchGitHubRepos(apiClient, l)
		if err != nil {
			log.Fatalf("Failed to search github repos: %v", err)
		}
		for _, v := range searchResp.Items {
			log.Printf("Getting tarbal for repo: %+v\n", v.FullName)
			resp, err := githubapi.GetRepoTarball(v.FullName)
			if err != nil {
				return
			}
			codeDir := "/tmp/repos"
			log.Printf("Extracting tarbal for repo: %+v\n", v.FullName)
			fileutils.ExtractTarball(resp, codeDir)
		}
	}
}
