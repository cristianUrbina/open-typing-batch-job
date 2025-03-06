package postgredatabase

import (
	"database/sql"
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"
)

type PostgresLanguageRepository struct {
	DB *sql.DB
}

// GetLanguages fetches all languages and their associated file extensions.
func (r *PostgresLanguageRepository) GetLanguages() ([]domain.Language, error) {
	// Query to join the language and language_file_ext tables
	query := `
		SELECT l.id, l.name, l.logo_url, e.extension
		FROM code.language l
		LEFT JOIN code.language_file_ext e ON l.id = e.language_id
		ORDER BY l.id
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query languages: %w", err)
	}
	defer rows.Close()

	// Map to store languages and their extensions
	languageMap := make(map[int]*domain.Language)

	for rows.Next() {
		var (
			id        int
			name      string
			logoURL   string
			extension sql.NullString // Use sql.NullString to handle NULL extensions
		)
		if err := rows.Scan(&id, &name, &logoURL, &extension); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Check if the language is already in the map
		if lang, exists := languageMap[id]; exists {
			// Append the extension if it exists
			if extension.Valid {
				lang.Extensions = append(lang.Extensions, extension.String)
			}
		} else {
			// Create a new language entry
			lang := &domain.Language{
				ID:         id,
				Name:       name,
				LogoURL:    logoURL,
				Extensions: []string{},
			}
			// Append the extension if it exists
			if extension.Valid {
				lang.Extensions = append(lang.Extensions, extension.String)
			}
			languageMap[id] = lang
		}
	}

	// Convert the map to a slice
	var languages []domain.Language
	for _, lang := range languageMap {
		languages = append(languages, *lang)
	}

	return languages, nil
}
