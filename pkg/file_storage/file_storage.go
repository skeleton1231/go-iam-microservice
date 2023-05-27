package storage

import (
	"fmt"
	"sync"

	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/options"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/file_storage/s3"
)

type FileStorage interface {
	Upload(file string) (string, error)
	Download(file string) (string, error)
	Delete(file string) error
	// 你可以在此处添加更多的文件存储方法
}

var (
	fileStorageInstance FileStorage
	once                sync.Once
)

// GetFileStorageFactoryOr create file storage factory with the given options.
func GetFileStorageFactoryOr(opts *options.FileStorageOptions) (FileStorage, error) {
	if opts == nil && fileStorageInstance == nil {
		return nil, fmt.Errorf("failed to get file storage factory")
	}

	var err error
	once.Do(func() {
		switch opts.Provider {
		case "s3":
			fileStorageInstance, err = s3.NewS3Storage(opts.S3Options.ToS3StorageConfig())
		// case "gcs":
		// 	fileStorageInstance, err = gcs.NewGCSStorage(opts.GCSOptions)
		default:
			err = fmt.Errorf("unknown file storage provider: %s", opts.Provider)
		}
	})

	if fileStorageInstance == nil || err != nil {
		return nil, fmt.Errorf("failed to get file storage factory, fileStorageInstance: %+v, error: %w", fileStorageInstance, err)
	}

	return fileStorageInstance, nil
}
