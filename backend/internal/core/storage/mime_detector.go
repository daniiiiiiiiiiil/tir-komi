package storage

import (
	"encoding/base64"
	"strings"
)

func DetectContentTypeByMagic(data []byte) string {
	if len(data) == 0 {
		return "application/octet-stream"
	}

	if len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return "image/jpeg"
	}

	if len(data) >= 8 &&
		data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 &&
		data[4] == 0x0D && data[5] == 0x0A && data[6] == 0x1A && data[7] == 0x0A {
		return "image/png"
	}

	if len(data) >= 6 &&
		(data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38 &&
			(data[4] == 0x37 || data[4] == 0x39) && data[5] == 0x61) {
		return "image/gif"
	}

	if len(data) >= 12 &&
		data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 &&
		data[8] == 0x57 && data[9] == 0x45 && data[10] == 0x42 && data[11] == 0x50 {
		return "image/webp"
	}

	if len(data) >= 2 && data[0] == 0x42 && data[1] == 0x4D {
		return "image/bmp"
	}

	if len(data) >= 5 && string(data[0:4]) == "%PDF" {
		return "application/pdf"
	}

	if IsBase64Image(string(data)) {
		return "image/jpeg" // default
	}

	return "application/octet-stream"
}

func IsBase64Image(s string) bool {
	prefixes := []string{
		"data:image/jpeg;base64,",
		"data:image/png;base64,",
		"data:image/gif;base64,",
		"data:image/webp;base64,",
		"data:image/bmp;base64,",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	if _, err := base64.StdEncoding.DecodeString(s); err == nil {
		return true
	}

	return false
}

func DecodeBase64Image(s string) ([]byte, string, error) {
	parts := strings.SplitN(s, ",", 2)
	var data string
	var contentType string

	if len(parts) == 2 && strings.HasPrefix(parts[0], "data:image/") {
		contentType = strings.TrimPrefix(parts[0], "data:")
		data = parts[1]
	} else {
		data = s
	}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, "", err
	}

	if contentType == "" {
		contentType = DetectContentTypeByMagic(decoded)
	}

	return decoded, contentType, nil
}
