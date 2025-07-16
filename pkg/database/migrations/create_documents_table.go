package migrations

import (
	"database/sql"
)

func CreateDocumentsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS documents (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			path VARCHAR(255) NOT NULL,
			uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Exec(query)
	return err
}
