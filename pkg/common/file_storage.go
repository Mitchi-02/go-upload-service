package common

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"strings"
	"upload-service/configs"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FileStorage struct {
	config configs.StorageConfig
	client *minio.Client
}

func NewFileStorage(config configs.StorageConfig) *FileStorage {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		panic(err)
	}

	return &FileStorage{config: config, client: client}
}

func (fs *FileStorage) StoreFile(file *multipart.FileHeader, filename string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	objectKey := fs.config.FolderName + "/" + filename + filepath.Ext(file.Filename)
	objectPath := fs.config.BucketName + "/" + objectKey
	ctx := context.Background()
	_, err = fs.client.PutObject(ctx, fs.config.BucketName, objectKey, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}
	return objectPath, nil
}

func (fs *FileStorage) DeleteFile(docPath string) error {
	ctx := context.Background()
	return fs.client.RemoveObject(ctx, fs.config.BucketName, strings.TrimPrefix(docPath, fs.config.BucketName+"/"), minio.RemoveObjectOptions{})
}

func (fs *FileStorage) GetFileURL(filename string) string {
	return fs.config.BaseURL + "/" + filename
}
