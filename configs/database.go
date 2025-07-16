package configs

import (
	"upload-service/configs/internal"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     internal.GetEnv("DB_HOST", "localhost"),
		Port:     internal.GetEnv("DB_PORT", "5432"),
		User:     internal.GetEnv("DB_USER", "upload-service"),
		Password: internal.GetEnv("DB_PASSWORD", "password"),
		DBName:   internal.GetEnv("DB_NAME", "main"),
	}
}
