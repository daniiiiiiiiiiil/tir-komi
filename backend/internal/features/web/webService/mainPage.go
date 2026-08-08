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
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot = "."
	}

	htmlPath := filepath.Join(projectRoot, "public", filename)

	html, err := s.webRepository.GetFile(htmlPath)
	if err != nil {
		return nil, fmt.Errorf("get file from repository: %w", err)
	}

	return html, nil
}
