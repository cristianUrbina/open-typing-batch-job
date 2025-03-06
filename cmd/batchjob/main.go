package main

import (
	"log"
	"os"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/githubrepo"

	infrastructure "cristianUrbina/open-typing-batch-job/internal/infrastructure/parser"
)

func main() {
	log.Printf("Searching for github repos")
	// languages := []string{"c", "python", "go", "js", "rust", "java"}
	languages := []string{"python", "js", "java"}
	for _, l := range languages {
		tmpDir, err := os.MkdirTemp("", "test-extract")
		if err != nil {
			log.Fatalf("failed to create temporary directory")
		}
		log.Printf("tmpDir %v", tmpDir)
		client := githubapiclient.NewAPIClient()
		githubRepo := githubrepo.NewRepositoryGithubRepo(*client)
		repoSearcher := domain.NewRepoSearcher(githubRepo)
		repoExtractor := domain.NewCodeExtractor(tmpDir)
		fileFilter := domain.NewFileFilter([]string{"c", "h", "python"})
		contentReader := domain.NewCodeFileContentReader()
		extr := infrastructure.NewTreeSitterSnippetExtractor()
		codeAnalyzer := domain.NewCodeAnalyzer(extr)
		contendDownloader := domain.NewRepositoryContentDownloader(githubRepo)

		repos, err := repoSearcher.SearchByLang(l)
		if err != nil {
			log.Fatalf("Failed searching repos %v", err)
		}
		for _, r := range repos {
			log.Printf("repo %v", r.Name)
			log.Println("getting  tarball")
			repoWithContent, err := contendDownloader.GetRepoContent(r)
			if err != nil {
				log.Printf("error getting repo content: %v", err)
				continue
			}
			// TODO: check why is causing extract to fail. Suspect it has to be with how the buffer is read when stored
			// fsRepo := &filesystemdb.FSRepositoryRepo{}
			// svc := app.NewCodeProjectService(fsRepo)
			// err = svc.AddRepo(repoWithContent)
			// if err != nil {
			// 	log.Fatalf("error creating code, %v", err)
			// }
			log.Println("extracting  tarball")
			files, err := repoExtractor.Extract(repoWithContent)
			if err != nil {
				log.Printf("error extracting tarball %v", err)
				continue
			}
			filteredFiles, err := fileFilter.Filter(files)
			if err != nil {
				log.Printf("error filtering files, %v", err)
				continue
			}
			for _, f := range filteredFiles {
				code, err := contentReader.Read(r, f)
				if err != nil {
					log.Printf("an unexpected error ocurred while getting file content, %v", err)
					continue
				}
				codeSnippets, err := codeAnalyzer.Analyze(code)
				if err != nil {
					log.Printf("error analyzing file content, %v", err)
					continue
				}
				for _, s := range codeSnippets {
					log.Printf("code snippet: %v", s.Content)
				}
			}
		}
	}
}
