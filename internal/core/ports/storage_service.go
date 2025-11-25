package ports

import "mime/multipart"

type StorageService interface {
	UploadFile(file *multipart.FileHeader, folder string) (string, error)
	DeleteFile(fileURL string) error
	GetFileURL(fileID string) string
}
