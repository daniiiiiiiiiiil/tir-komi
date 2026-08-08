package storage

import (
	"strings"
)

func DetectContentType(data []byte) string {
	if len(data) == 0 {
		return "application/octet-stream"
	}

	if len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return "image/jpeg"
	}
	if len(data) >= 8 && data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 &&
		data[4] == 0x0D && data[5] == 0x0A && data[6] == 0x1A && data[7] == 0x0A {
		return "image/png"
	}
	if len(data) >= 6 && data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38 &&
		(data[4] == 0x37 || data[4] == 0x39) && data[5] == 0x61 {
		return "image/gif"
	}
	if len(data) >= 5 && string(data[0:4]) == "%PDF" {
		return "application/pdf"
	}

	return "application/octet-stream"
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
