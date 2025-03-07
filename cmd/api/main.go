package main

import (
	"log"
	"net/http"

	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/postgredatabase"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/httphandlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type LanguageDto struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Extensions []string `json:"extensions"`
	LogoURL    string   `json:"logo_url"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
	db, err := postgredatabase.NewDatabase()
	if err != nil {
		log.Fatalf("error creating db connection: %v", err)
	}
	defer db.Close()

	repo := &postgredatabase.PostgresLanguageRepository{DB: db}
	service := &app.LanguageService{Repo: repo}
	handler := &httphandlers.LanguageHandler{Service: service}

	http.HandleFunc("/languages", handler.GetLanguages)

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
