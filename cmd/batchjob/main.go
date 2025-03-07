package main

import (
	"log"
	"os"
	"runtime"
	"sync"

	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/dynamodb"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/githubrepo"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/postgredatabase"

	infrastructure "cristianUrbina/open-typing-batch-job/internal/infrastructure/parser"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}

	log.Println("Searching for GitHub repos...")

	db, err := postgredatabase.NewDatabase()
	if err != nil {
		log.Fatalf("Error creating DB connection: %v", err)
	}
	repo := &postgredatabase.PostgresLanguageRepository{DB: db}
	langSvc := &app.LanguageService{Repo: repo}

	langs, err := langSvc.GetAvailableLanguages()
	if err != nil {
		log.Fatalf("Error getting languages: %v", err)
	}

	dynamoClient := dynamodb.NewDynamoClient()
	snippetRepo := dynamodb.NewCodeSnippetRepository(dynamoClient)
	snippetSvc := app.NewSnippetService(snippetRepo)

	var wg sync.WaitGroup
	repoChan := make(chan domain.Repository, 50)

	for _, lang := range langs {
		wg.Add(1)
		go processLanguage(lang, &wg, repoChan)
	}

	go func() {
		wg.Wait()
		close(repoChan)
	}()

	numWorkers := runtime.NumCPU()
	var resultsWg sync.WaitGroup
	for range numWorkers {
		resultsWg.Add(1)
		go processRepoResults(repoChan, &resultsWg, snippetSvc)
	}

	resultsWg.Wait()

	log.Println("Processing completed!")
}

func processLanguage(lang domain.Language, wg *sync.WaitGroup, repoChan chan<- domain.Repository) {
	defer wg.Done()

	tmpDir, err := os.MkdirTemp("", "test-extract")
	if err != nil {
		log.Printf("Failed to create temp dir: %v", err)
		return
	}

	log.Printf("Processing language: %s in tempDir: %v", lang.Alias, tmpDir)

	client := githubapiclient.NewAPIClient()
	githubRepo := githubrepo.NewRepositoryGithubRepo(*client)
	repoSvc := app.NewRepoService(githubRepo)

	repos, err := repoSvc.SearchByLang(&lang)
	if err != nil {
		log.Printf("Failed searching repos for %s: %v", lang.Alias, err)
		return
	}

	for _, repo := range repos {
		repoChan <- repo
	}
}

func processRepoResults(repoChan <-chan domain.Repository, resultsWg *sync.WaitGroup, snippetSvc *app.SnippetService) {
	defer resultsWg.Done()

	client := githubapiclient.NewAPIClient()
	githubRepo := githubrepo.NewRepositoryGithubRepo(*client)
	contentReader := domain.NewCodeFileContentReader()
	extr := infrastructure.NewTreeSitterSnippetExtractor()
	codeSvc := app.NewCodeService(extr)
	repoSvc := app.NewRepoService(githubRepo)

	for repo := range repoChan {
		snippetsCnt := 0
		fileFilter := domain.NewFileFilter(repo.Lang.Extensions)
		log.Printf("Getting content: %s", repo.Name)

		repoWithContent, err := repoSvc.GetRepoContent(repo)
		if err != nil {
			log.Printf("Error getting repo content: %v", err)
			continue
		}

		log.Printf("Extracting: %s", repo.Name)
		files, err := repoSvc.Extract(repoWithContent)
		if err != nil {
			log.Printf("Error extracting tarball: %v", err)
			continue
		}

		log.Printf("Filtering files: %s", repo.Name)
		filteredFiles, err := fileFilter.Filter(files)
		if err != nil {
			log.Printf("Error filtering files: %v", err)
			continue
		}

		log.Printf("Analyzing files: %s", repo.Name)
		for _, file := range filteredFiles {
			code, err := contentReader.Read(repo, file)
			if err != nil {
				log.Printf("Error reading file content: %v", err)
				continue
			}

			codeSnippets, err := codeSvc.Analyze(code)
			if err != nil {
				log.Printf("Error analyzing file content: %v", err)
				continue
			}

			snippetsCnt = snippetsCnt + len(codeSnippets)
			for _, snippet := range codeSnippets {
				err := snippetSvc.AddSnippet(repo.Name, file, snippet.Language, snippet.Content)
				if err != nil {
					log.Printf("Error persisting snippet to DynamoDB: %v", err)
				}
			}
		}
		log.Printf("no of snippets extracted from %v, %v", repo.Name, snippetsCnt)
	}
}
