package users

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByEmail(email string) (User, error) {
	var user User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func (r *Repository) Create(user User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := r.db.Exec(query, user.Email, user.Password)
	return err
}
