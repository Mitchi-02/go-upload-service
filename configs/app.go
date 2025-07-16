package configs

import "upload-service/configs/internal"

type AppConfig struct {
	Port      string
	JWTSecret string
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Port:      internal.GetEnv("APP_PORT", "8080"),
		JWTSecret: internal.GetEnv("JWT_SECRET", "lmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
}
