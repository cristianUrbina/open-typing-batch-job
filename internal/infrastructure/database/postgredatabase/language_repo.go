package postgredatabase

import (
	"database/sql"
	"fmt"

	"cristianUrbina/open-typing-batch-job/internal/domain"

	"github.com/lib/pq"
)

type PostgresLanguageRepository struct {
	DB *sql.DB
}

func (r *PostgresLanguageRepository) GetLanguages() ([]domain.Language, error) {
	query := `
        SELECT
            l.id,
            l.name,
            l.alias,
            l.logo_url,
            COALESCE(ARRAY_AGG(DISTINCT e.extension) FILTER (WHERE e.extension IS NOT NULL), '{}') AS extensions,
            COALESCE(ARRAY_AGG(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL), '{}') AS capabilities
        FROM code.language l
        LEFT JOIN code.language_file_ext e ON l.id = e.language_id
        LEFT JOIN code.language_capability lc ON l.id = lc.language_id
        LEFT JOIN code.capability c ON lc.capability_id = c.id
        GROUP BY l.id, l.name, l.logo_url
        ORDER BY l.id;
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query languages: %w", err)
	}
	defer rows.Close()

	var languages []domain.Language

	for rows.Next() {
		var lang domain.Language
		var extensions pq.StringArray
		var capabilities pq.StringArray

		if err := rows.Scan(&lang.ID, &lang.Name, &lang.Alias, &lang.LogoURL, &extensions, &capabilities); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		lang.Extensions = extensions
		lang.Capabilities = capabilities
		languages = append(languages, lang)
	}

	return languages, nil
}
