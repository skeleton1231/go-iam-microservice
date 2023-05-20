package storage

type FileStorage interface {
	UploadFile(file string) (string, error)
	DeleteFile(file string) error
}
