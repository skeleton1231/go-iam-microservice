// s3/s3.go
package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	Bucket  string
	Session *session.Session
}

type S3StorageConfig struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
}

func (c *S3StorageConfig) GetStorageType() string {
	return "s3"
}

func NewS3Storage(config *S3StorageConfig) (*S3Storage, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID, config.SecretAccessKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return &S3Storage{
		Bucket:  config.Bucket,
		Session: sess,
	}, nil
}

// Implement other methods for S3Storage

func (s *S3Storage) Upload(file string) (string, error) {
	uploader := s3manager.NewUploader(s.Session)

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
		Body:   f,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func (s *S3Storage) Download(file string) (string, error) {
	downloader := s3manager.NewDownloader(s.Session)

	f, err := os.Create(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})

	if err != nil {
		return "", err
	}

	// Return local path to the downloaded file
	return f.Name(), nil
}

func (s *S3Storage) Delete(file string) error {
	svc := s3.New(s.Session)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})

	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})

	return err
}
