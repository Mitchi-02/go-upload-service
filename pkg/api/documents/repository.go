package documents

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func DocumentRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetDocumentsPaginated(limit, offset int) ([]Document, error) {
	rows, err := r.db.Query("SELECT id, name, path, uploaded_at FROM documents ORDER BY uploaded_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	documents := make([]Document, 0)
	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.ID, &doc.Name, &doc.Path, &doc.UploadedAt); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (r *Repository) GetTotalDocuments() (int, error) {
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM documents").Scan(&total)
	return total, err
}

func (r *Repository) Store(doc Document) error {
	query := `INSERT INTO documents (id, name, path, uploaded_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, doc.ID, doc.Name, doc.Path, doc.UploadedAt)
	return err
}

func (r *Repository) GetByID(id string) (Document, error) {
	var doc Document
	query := `SELECT id, name, path, uploaded_at FROM documents WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&doc.ID, &doc.Name, &doc.Path, &doc.UploadedAt)
	return doc, err
}

func (r *Repository) Delete(id string) error {
	query := `DELETE FROM documents WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
