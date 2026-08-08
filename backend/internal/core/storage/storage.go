package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type FileStorage struct {
	uploadDir string
	maxSize   int64
}

type FileStorageConfig struct {
	UploadDir string
	MaxSize   int64
}

func NewFileStorage(config FileStorageConfig) *FileStorage {
	if config.MaxSize == 0 {
		config.MaxSize = 10 << 20
	}
	if config.UploadDir == "" {
		config.UploadDir = "./uploads"
	}

	if err := os.MkdirAll(config.UploadDir, 0755); err != nil {
		panic(fmt.Errorf("failed to create upload directory: %w", err))
	}

	return &FileStorage{
		uploadDir: config.UploadDir,
		maxSize:   config.MaxSize,
	}
}

func (s *FileStorage) Save(ctx context.Context, data []byte) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("empty file data")
	}

	if int64(len(data)) > s.maxSize {
		return "", fmt.Errorf("file too large: max %d bytes", s.maxSize)
	}

	contentType := DetectContentType(data)
	ext := GetExtension(contentType)

	if IsImage(contentType) {
		compressed, err := s.CompressImage(data)
		if err != nil {
			return "", fmt.Errorf("failed to compress image: %w", err)
		}
		data = compressed
	}

	filename := s.generateFilename(ext)
	filePath := filepath.Join(s.uploadDir, filename)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filePath, nil
}

func (s *FileStorage) SaveMultipart(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader.Size > s.maxSize {
		return "", fmt.Errorf("file too large: max %d bytes", s.maxSize)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return s.Save(ctx, data)
}

func (s *FileStorage) Delete(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

func (s *FileStorage) generateFilename(ext string) string {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	uuid := uuid.New().String()[:8]
	return fmt.Sprintf("%s_%s%s", timestamp, uuid, ext)
}

func (s *FileStorage) GetFullURL(filePath string, baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, filepath.Base(filePath))
}
