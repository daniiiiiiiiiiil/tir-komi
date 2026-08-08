package storage

import (
	"mime"
	"path/filepath"
	"strings"
)

func DetectContentType(data []byte) string {
	if len(data) == 0 {
		return "application/octet-stream"
	}
	return mime.TypeByExtension(filepath.Ext(""))
}

func GetExtension(contentType string) string {
	switch {
	case IsImage(contentType):
		switch contentType {
		case "image/jpeg":
			return ".jpg"
		case "image/png":
			return ".png"
		case "image/gif":
			return ".gif"
		case "image/webp":
			return ".webp"
		default:
			return ".jpg"
		}
	case IsPDF(contentType):
		return ".pdf"
	default:
		return ".bin"
	}
}

func IsImage(contentType string) bool {
	imageTypes := map[string]bool{
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/webp":    true,
		"image/bmp":     true,
		"image/svg+xml": true,
	}
	return imageTypes[contentType]
}

func IsPDF(contentType string) bool {
	return contentType == "application/pdf"
}

func IsPDFBytes(data []byte) bool {
	if len(data) < 5 {
		return false
	}
	return string(data[0:4]) == "%PDF"
}

func SanitizeFilename(filename string) string {
	filename = strings.ReplaceAll(filename, " ", "_")
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, ":", "_")
	filename = strings.ReplaceAll(filename, "*", "_")
	filename = strings.ReplaceAll(filename, "?", "_")
	filename = strings.ReplaceAll(filename, "\"", "_")
	filename = strings.ReplaceAll(filename, "<", "_")
	filename = strings.ReplaceAll(filename, ">", "_")
	filename = strings.ReplaceAll(filename, "|", "_")
	return filename
}
