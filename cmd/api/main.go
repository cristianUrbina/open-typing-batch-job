package main

import (
	"log"
	"net/http"

	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/dynamodb"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/postgredatabase"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/httphandlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type LanguageDto struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Extensions []string `json:"extensions"`
	LogoURL    string   `json:"logo_url"`
}

func main() {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	db, err := postgredatabase.NewDatabase()
	if err != nil {
		log.Fatalf("error creating db connection: %v", err)
	}
	defer db.Close()

	languageRepo := &postgredatabase.PostgresLanguageRepository{DB: db}
	langSvc := &app.LanguageService{Repo: languageRepo}
	langHandler := &httphandlers.LanguageHandler{Service: langSvc}

	dynamoClient := dynamodb.NewDynamoClient()
	snippetRepo := dynamodb.NewCodeSnippetRepository(dynamoClient)
	snippetSvc := app.NewSnippetService(snippetRepo)
	snippetHandler := &httphandlers.SnippetHandler{Service: snippetSvc}

	r := mux.NewRouter()
	r.HandleFunc("/languages", langHandler.GetLanguages).Methods("GET")
	r.HandleFunc("/languages/{lang}", langHandler.GetLanguageByName).Methods("GET")
	r.HandleFunc("/snippets/{lang}", snippetHandler.GetSnippetByLanguage).Methods("GET")
	r.HandleFunc("/version", httphandlers.GetVersion).Methods("GET")

	log.Println("Server is running on port 8080...")
	handlerWithCors := corsHandler.Handler(r)

	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		log.Fatal(err)
	}
}
