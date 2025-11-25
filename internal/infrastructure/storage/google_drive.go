package storage

import (
	"context"
	"ecommerce-backend/internal/core/ports"
	"ecommerce-backend/internal/infrastructure/config"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type googleDriveStorage struct {
	service  *drive.Service
	folderID string
}

func NewGoogleDriveStorage(cfg *config.Config) (ports.StorageService, error) {
	ctx := context.Background()
	
	service, err := drive.NewService(ctx, option.WithCredentialsFile(cfg.GoogleDrive.CredentialsPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create drive service: %v", err)
	}

	return &googleDriveStorage{
		service:  service,
		folderID: cfg.GoogleDrive.FolderID,
	}, nil
}

func (s *googleDriveStorage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	driveFile := &drive.File{
		Name:    filename,
		Parents: []string{s.folderID},
	}

	res, err := s.service.Files.Create(driveFile).Media(src).Do()
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	_, err = s.service.Permissions.Create(res.Id, &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}).Do()
	if err != nil {
		return "", fmt.Errorf("failed to set file permissions: %v", err)
	}

	fileURL := fmt.Sprintf("https://drive.google.com/uc?id=%s", res.Id)
	return fileURL, nil
}

func (s *googleDriveStorage) DeleteFile(fileURL string) error {
	fileID := extractFileIDFromURL(fileURL)
	if fileID == "" {
		return fmt.Errorf("invalid file URL")
	}

	err := s.service.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

func (s *googleDriveStorage) GetFileURL(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/uc?id=%s", fileID)
}

func extractFileIDFromURL(fileURL string) string {
	var fileID string
	fmt.Sscanf(fileURL, "https://drive.google.com/uc?id=%s", &fileID)
	return fileID
}

type localStorage struct {
	uploadPath string
	baseURL    string
}

func NewLocalStorage(uploadPath, baseURL string) ports.StorageService {
	return &localStorage{
		uploadPath: uploadPath,
		baseURL:    baseURL,
	}
}

func (s *localStorage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	fileURL := fmt.Sprintf("%s/%s/%s", s.baseURL, folder, filename)
	
	return fileURL, nil
}

func (s *localStorage) DeleteFile(fileURL string) error {
	return nil
}

func (s *localStorage) GetFileURL(fileID string) string {
	return fmt.Sprintf("%s/%s", s.baseURL, fileID)
}
