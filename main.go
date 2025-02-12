package main

import (
	"log"

	"cristianUrbina/open-typing-batch-job/fileutils"
)

func main() {
	log.Printf("Searching for github repos")
	languages := []string{"c", "python", "go", "js", "rust", "java"}
	for _, l := range languages {
		searchResp := searchGitHubRepos(l)
		for _, v := range searchResp.Items {
			log.Printf("Getting tarbal for repo: %+v\n", v.FullName)
			resp, err := getRepoTarball(v.FullName)
			if err != nil {
				return
			}
			codeDir := "./batch_job_pulled_repos"
			log.Printf("Extracting tarbal for repo: %+v\n", v.FullName)
			fileutils.ExtractTarball(resp, codeDir)
		}
	}
}
