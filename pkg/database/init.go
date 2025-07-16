package database

import (
	"database/sql"
	"fmt"
	"upload-service/configs"
	"upload-service/pkg/database/migrations"

	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
	config := configs.GetDatabaseConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := migrations.CreateDocumentsTable(db); err != nil {
		return nil, err
	}

	if err := migrations.CreateUsersTable(db); err != nil {
		return nil, err
	}

	return db, nil
}
