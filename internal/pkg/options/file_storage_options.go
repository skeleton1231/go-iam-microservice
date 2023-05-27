package options

import (
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/file_storage/s3"
	"github.com/spf13/pflag"
)

// FileStorageOptions defines options for file storage (like S3, GCS, etc.).
type FileStorageOptions struct {
	Provider   string      `json:"provider"   mapstructure:"provider"   description:"File storage service provider"`
	S3Options  *S3Options  `json:"s3"         mapstructure:"s3"         description:"Configuration options for Amazon S3"`
	GCSOptions *GCSOptions `json:"gcs"      mapstructure:"gcs"        description:"Configuration options for Google Cloud Storage"`
	// Add other storage options here
}

// S3Options defines configuration options for Amazon S3.
type S3Options struct {
	AccessKeyID     string `json:"accessKeyId"     mapstructure:"accessKeyId"     description:"Access key ID for Amazon S3"`
	SecretAccessKey string `json:"secretAccessKey" mapstructure:"secretAccessKey" description:"Secret access key for Amazon S3"`
	Region          string `json:"region"          mapstructure:"region"          description:"Region for Amazon S3"`
	Bucket          string `json:"bucket"          mapstructure:"bucket"          description:"Bucket name for Amazon S3"`
}

// ToS3StorageConfig converts S3Options to s3.S3StorageConfig.
func (opts *S3Options) ToS3StorageConfig() *s3.S3StorageConfig {
	if opts == nil {
		return nil
	}

	return &s3.S3StorageConfig{
		AccessKeyID:     opts.AccessKeyID,
		SecretAccessKey: opts.SecretAccessKey,
		Region:          opts.Region,
		Bucket:          opts.Bucket,
	}
}

// GCSOptions defines configuration options for Google Cloud Storage.
type GCSOptions struct {
	ServiceAccountKeyFile string `json:"serviceAccountKeyFile" mapstructure:"serviceAccountKeyFile" description:"Service account key file for Google Cloud Storage"`
	Bucket                string `json:"bucket"                mapstructure:"bucket"                description:"Bucket name for Google Cloud Storage"`
}

// NewFileStorageOptions creates a 'zero' value instance.
func NewFileStorageOptions() *FileStorageOptions {
	return &FileStorageOptions{
		Provider: "S3",
		S3Options: &S3Options{
			AccessKeyID:     "",
			SecretAccessKey: "",
			Region:          "",
			Bucket:          "",
		},
		GCSOptions: &GCSOptions{
			ServiceAccountKeyFile: "",
			Bucket:                "",
		},
		// Initialize other storage options here
	}
}

// Validate verifies flags passed to FileStorageOptions.
func (o *FileStorageOptions) Validate() []error {
	errs := []error{}

	// Add validation rules here

	return errs
}

// AddFlags adds flags related to file storage for a specific APIServer to the specified FlagSet.
func (o *FileStorageOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Provider, "file-storage.provider", o.Provider, "File storage service provider.")

	fs.StringVar(&o.S3Options.AccessKeyID, "file-storage.s3.accessKeyId", o.S3Options.AccessKeyID, "Access key ID for Amazon S3.")
	fs.StringVar(&o.S3Options.SecretAccessKey, "file-storage.s3.secretAccessKey", o.S3Options.SecretAccessKey, "Secret access key for Amazon S3.")
	fs.StringVar(&o.S3Options.Region, "file-storage.s3.region", o.S3Options.Region, "Region for Amazon S3.")
	fs.StringVar(&o.S3Options.Bucket, "file-storage.s3.bucket", o.S3Options.Bucket, "Bucket name for Amazon S3.")

	fs.StringVar(&o.GCSOptions.ServiceAccountKeyFile, "file-storage.gcs.serviceAccountKeyFile", o.GCSOptions.ServiceAccountKeyFile, "Service account key file for Google Cloud Storage.")
	fs.StringVar(&o.GCSOptions.Bucket, "file-storage.gcs.bucket", o.GCSOptions.Bucket, "Bucket name for Google Cloud Storage.")

	// Add flags for other storage options here
}
