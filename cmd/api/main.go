package main

import (
	"cristianUrbina/open-typing-batch-job/internal/app"
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/internal/infrastructure/database/postgredatabase"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// log.Printf("dbname %v", dbName)
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	repo := &postgredatabase.PostgresLanguageRepository{DB: db}
	service := &domain.LanguageService{Repo: repo}
	handler := &app.LanguageHandler{Service: service}

	// Set up the HTTP route
	http.HandleFunc("/languages", handler.GetLanguages)

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
