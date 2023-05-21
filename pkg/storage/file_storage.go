package storage

// storage/storage.go

import (
	"fmt"

	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/storage/s3"
)

// StorageConfig is an interface representing the configuration for different storage implementations.
type StorageConfig interface {
	GetStorageType() string
}

// FileStorage is an interface representing the required methods for file storage.
type FileStorage interface {
	Upload(file string) (string, error)
	Download(file string) (string, error)
	Delete(file string) error
}

func StorageFactory(config StorageConfig) (FileStorage, error) {
	switch config.GetStorageType() {
	case "s3":
		return s3.NewS3Storage(config.(*s3.S3StorageConfig))
	default:
		return nil, fmt.Errorf("unknown storage type: %s", config.GetStorageType())
	}
}
