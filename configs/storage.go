package configs

import (
	"upload-service/configs/internal"
)

type StorageConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	FolderName      string
	BaseURL         string
}

func GetStorageConfig() StorageConfig {
	return StorageConfig{
		Endpoint:        internal.GetEnv("MINIO_ENDPOINT", "localhost:9000"),
		AccessKeyID:     internal.GetEnv("MINIO_ACCESS_KEY", "L6G55UE1961PUKFENWFD"),
		SecretAccessKey: internal.GetEnv("MINIO_SECRET_KEY", "RCmF4UM9wP6QMqddyixWqZv9gUKSZS76MHMXAwtz"),
		BucketName:      internal.GetEnv("MINIO_BUCKET_NAME", "upload-service"),
		FolderName:      internal.GetEnv("MINIO_FOLDER_NAME", "documents"),
		BaseURL:         internal.GetEnv("MINIO_BASE_URL", "http://localhost:9000"),
	}
}
