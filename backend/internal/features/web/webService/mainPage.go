// internal/features/web/webService/mainPage.go

package webService

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *WebService) GetMainPage() ([]byte, error) {
	return s.GetHTMLFile("/index.html")
}

func (s *WebService) GetAdminPage() ([]byte, error) {
	return s.GetHTMLFile("/admin.html")
}

func (s *WebService) GetHTMLFile(filename string) ([]byte, error) {
	paths := []string{
		filepath.Join("backend", "public", filename),
		filepath.Join(".", "public", filename),
	}

	if projectRoot := os.Getenv("PROJECT_ROOT"); projectRoot != "" {
		paths = append([]string{filepath.Join(projectRoot, "public", filename)}, paths...)
	}

	var lastErr error
	for _, htmlPath := range paths {
		html, err := s.webRepository.GetFile(htmlPath)
		if err == nil {
			return html, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("get file from repository: %w", lastErr)
}
